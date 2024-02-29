package logpoller

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
)

// ORM represents the persistent data access layer used by the log poller. At this moment, it's a bit leaky abstraction, because
// it exposes some of the database implementation details (e.g. pg.Q). Ideally it should be agnostic and could be applied to any persistence layer.
// What is more, LogPoller should not be aware of the underlying database implementation and delegate all the queries to the ORM.
type ORM interface {
	InsertLogs(ctx context.Context, logs []Log) error
	InsertLogsWithBlock(ctx context.Context, logs []Log, block LogPollerBlock) error
	InsertFilter(ctx context.Context, filter Filter) error

	LoadFilters(ctx context.Context) (map[string]Filter, error)
	DeleteFilter(ctx context.Context, name string) error

	InsertBlock(ctx context.Context, blockHash common.Hash, blockNumber int64, blockTimestamp time.Time, finalizedBlock int64) error
	DeleteBlocksBefore(ctx context.Context, end int64, limit int64) (int64, error)
	DeleteLogsAndBlocksAfter(ctx context.Context, start int64) error
	DeleteExpiredLogs(ctx context.Context, limit int64) (int64, error)

	GetBlocksRange(ctx context.Context, start int64, end int64) ([]LogPollerBlock, error)
	SelectBlockByNumber(ctx context.Context, blockNumber int64) (*LogPollerBlock, error)
	SelectBlockByHash(ctx context.Context, hash common.Hash) (*LogPollerBlock, error)
	SelectLatestBlock(ctx context.Context) (*LogPollerBlock, error)

	SelectLogs(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash) ([]Log, error)
	SelectLogsWithSigs(ctx context.Context, start, end int64, address common.Address, eventSigs []common.Hash) ([]Log, error)
	SelectLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, after time.Time, confs Confirmations) ([]Log, error)
	SelectLatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs Confirmations) (*Log, error)
	SelectLatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, addresses []common.Address, eventSigs []common.Hash, confs Confirmations) ([]Log, error)
	SelectLatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs Confirmations) (int64, error)
	SelectLogsByBlockRange(ctx context.Context, start, end int64) ([]Log, error)

	SelectIndexedLogs(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, confs Confirmations) ([]Log, error)
	SelectIndexedLogsByBlockRange(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash) ([]Log, error)
	SelectIndexedLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, after time.Time, confs Confirmations) ([]Log, error)
	SelectIndexedLogsTopicGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin common.Hash, confs Confirmations) ([]Log, error)
	SelectIndexedLogsTopicRange(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin, topicValueMax common.Hash, confs Confirmations) ([]Log, error)
	SelectIndexedLogsWithSigsExcluding(ctx context.Context, sigA, sigB common.Hash, topicIndex int, address common.Address, startBlock, endBlock int64, confs Confirmations) ([]Log, error)
	SelectIndexedLogsByTxHash(ctx context.Context, address common.Address, eventSig common.Hash, txHash common.Hash) ([]Log, error)
	SelectLogsDataWordRange(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin, wordValueMax common.Hash, confs Confirmations) ([]Log, error)
	SelectLogsDataWordGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin common.Hash, confs Confirmations) ([]Log, error)
	SelectLogsDataWordBetween(ctx context.Context, address common.Address, eventSig common.Hash, wordIndexMin int, wordIndexMax int, wordValue common.Hash, confs Confirmations) ([]Log, error)
}

type orm struct {
	chainID *big.Int
	db      sqlutil.Queryer
	lggr    logger.Logger
}

var _ ORM = &orm{}

// NewORM creates an orm scoped to chainID.
func NewORM(chainID *big.Int, db sqlutil.Queryer, lggr logger.Logger) ORM {
	return &orm{
		chainID: chainID,
		db:      db,
		lggr:    lggr,
	}
}

func (o *orm) Transaction(ctx context.Context, fn func(*orm) error) (err error) {
	return sqlutil.Transact(ctx, o.new, o.db, nil, fn)
}

// new returns a NewORM like o, but backed by q.
func (o *orm) new(q sqlutil.Queryer) *orm { return NewORM(o.chainID, q, o.lggr).(*orm) }

// InsertBlock is idempotent to support replays.
func (o *orm) InsertBlock(ctx context.Context, blockHash common.Hash, blockNumber int64, blockTimestamp time.Time, finalizedBlock int64) error {
	args, err := newQueryArgs(o.chainID).
		withCustomHashArg("block_hash", blockHash).
		withCustomArg("block_number", blockNumber).
		withCustomArg("block_timestamp", blockTimestamp).
		withCustomArg("finalized_block_number", finalizedBlock).
		toArgs()
	if err != nil {
		return err
	}
	query := `INSERT INTO evm.log_poller_blocks 
				(evm_chain_id, block_hash, block_number, block_timestamp, finalized_block_number, created_at) 
      		VALUES (:evm_chain_id, :block_hash, :block_number, :block_timestamp, :finalized_block_number, NOW()) 
			ON CONFLICT DO NOTHING`
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	_, err = o.db.ExecContext(ctx, query, sqlArgs...)
	return err
}

// InsertFilter is idempotent.
//
// Each address/event pair must have a unique job id, so it may be removed when the job is deleted.
// If a second job tries to overwrite the same pair, this should fail.
func (o *orm) InsertFilter(ctx context.Context, filter Filter) (err error) {
	// '::' has to be escaped in the query string
	// https://github.com/jmoiron/sqlx/issues/91, https://github.com/jmoiron/sqlx/issues/428
	/*
		query := `
			INSERT INTO evm.log_poller_filters
		  		(name, evm_chain_id, retention, created_at, address, event)
			SELECT * FROM
				(SELECT $1, $2 ::NUMERIC, $3 ::BIGINT, NOW()) x,
				(SELECT unnest($4 ::BYTEA[]) addr) a,
				(SELECT unnest($5 ::BYTEA[]) ev) e
			ON CONFLICT (evm.f_log_poller_filter_hash(name, evm_chain_id, address, event, topic2, topic3, topic4))
			DO UPDATE SET retention=$3 ::BIGINT, max_logs_kept=$6 ::NUMERIC, logs_per_block=$7 ::NUMERIC`
	*/

	topicArrays := []types.HashArray{filter.Topic2, filter.Topic3, filter.Topic4}
	args, err := newQueryArgs(o.chainID).
		withCustomArg("name", filter.Name).
		withRetention(filter.Retention).
		withMaxLogsKept(filter.MaxLogsKept).
		withLogsPerBlock(filter.LogsPerBlock).
		withAddressArray(filter.Addresses).
		withEventSigArray(filter.EventSigs).
		withTopicArrays(filter.Topic2, filter.Topic3, filter.Topic4).
		toArgs()
	if err != nil {
		return err
	}
	// '::' has to be escaped in the query string
	// https://github.com/jmoiron/sqlx/issues/91, https://github.com/jmoiron/sqlx/issues/428
	var topicsColumns, topicsSql strings.Builder
	for n, topicValues := range topicArrays {
		if len(topicValues) != 0 {
			topicCol := fmt.Sprintf("topic%d", n+2)
			fmt.Fprintf(&topicsColumns, ", %s", topicCol)
			fmt.Fprintf(&topicsSql, ",\n(SELECT unnest(:%s ::::BYTEA[]) %s) t%d", topicCol, topicCol, n+2)
		}
	}
	query := fmt.Sprintf(`
		INSERT INTO evm.log_poller_filters
	  		(name, evm_chain_id, retention, max_logs_kept, logs_per_block, created_at, address, event %s)
		SELECT * FROM
			(SELECT :name, :evm_chain_id ::::NUMERIC, :retention ::::BIGINT, :max_logs_kept ::::NUMERIC, :logs_per_block ::::NUMERIC, NOW()) x,
			(SELECT unnest(:address_array ::::BYTEA[]) addr) a,
			(SELECT unnest(:event_sig_array ::::BYTEA[]) ev) e
			%s
		ON CONFLICT  (evm.f_log_poller_filter_hash(name, evm_chain_id, address, event, topic2, topic3, topic4))
		DO UPDATE SET retention=:retention ::::BIGINT, max_logs_kept=:max_logs_kept ::::NUMERIC, logs_per_block=:logs_per_block ::::NUMERIC`,
		topicsColumns.String(),
		topicsSql.String())

	query, sqlArgs, _ := o.db.BindNamed(query, args)
	_, err = o.db.ExecContext(ctx, query, sqlArgs...)
	return err
}

// DeleteFilter removes all events,address pairs associated with the Filter
func (o *orm) DeleteFilter(ctx context.Context, name string) error {
	_, err := o.db.ExecContext(ctx,
		`DELETE FROM evm.log_poller_filters WHERE name = $1 AND evm_chain_id = $2`,
		name, ubig.New(o.chainID))
	return err

}

// LoadFilters returns all filters for this chain
func (o *orm) LoadFilters(ctx context.Context) (map[string]Filter, error) {
	query := `SELECT name,
			ARRAY_AGG(DISTINCT address)::BYTEA[] AS addresses, 
			ARRAY_AGG(DISTINCT event)::BYTEA[] AS event_sigs,
			ARRAY_AGG(DISTINCT topic2 ORDER BY topic2) FILTER(WHERE topic2 IS NOT NULL) AS topic2,
			ARRAY_AGG(DISTINCT topic3 ORDER BY topic3) FILTER(WHERE topic3 IS NOT NULL) AS topic3,
			ARRAY_AGG(DISTINCT topic4 ORDER BY topic4) FILTER(WHERE topic4 IS NOT NULL) AS topic4,
			MAX(logs_per_block) AS logs_per_block,
			MAX(retention) AS retention,
			MAX(max_logs_kept) AS max_logs_kept
		FROM evm.log_poller_filters WHERE evm_chain_id = $1
		GROUP BY name`
	var rows []Filter
	err := o.db.SelectContext(ctx, &rows, query, ubig.New(o.chainID))
	filters := make(map[string]Filter)
	for _, filter := range rows {
		filters[filter.Name] = filter
	}
	return filters, err
}

func (o *orm) SelectBlockByHash(ctx context.Context, hash common.Hash) (*LogPollerBlock, error) {
	var b LogPollerBlock
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE block_hash = $1 AND evm_chain_id = $2`, hash.Bytes(), ubig.New(o.chainID)); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *orm) SelectBlockByNumber(ctx context.Context, n int64) (*LogPollerBlock, error) {
	var b LogPollerBlock
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE block_number = $1 AND evm_chain_id = $2`, n, ubig.New(o.chainID)); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *orm) SelectLatestBlock(ctx context.Context) (*LogPollerBlock, error) {
	var b LogPollerBlock
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE evm_chain_id = $1 ORDER BY block_number DESC LIMIT 1`, ubig.New(o.chainID)); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *orm) SelectLatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs Confirmations) (*Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs
			WHERE evm_chain_id = :evm_chain_id
			AND event_sig = :event_sig
			AND address = :address
			AND block_number <= %s
			ORDER BY (block_number, log_index) DESC LIMIT 1`, nestedBlockNumberQuery(confs))
	var l Log

	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.GetContext(ctx, &l, query, sqlArgs...); err != nil {
		return nil, err
	}
	return &l, nil
}

// DeleteBlocksBefore delete all blocks before and including end.
func (o *orm) DeleteBlocksBefore(ctx context.Context, end int64, limit int64) (int64, error) {
	if limit > 0 {
		fmt.Println("Deleting all blocks before with limit", end, limit)
		result, err := o.db.ExecContext(ctx,
			`DELETE FROM evm.log_poller_blocks
        				WHERE block_number IN (
            				SELECT block_number FROM evm.log_poller_blocks
            				WHERE block_number <= $1 
            				AND evm_chain_id = $2
							LIMIT $3
						)
						AND evm_chain_id = $2`,
			end, ubig.New(o.chainID), limit)
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, err
	}
	fmt.Println("Deleting all blocks before", end)
	result, err := o.db.ExecContext(ctx, `DELETE FROM evm.log_poller_blocks 
       WHERE block_number <= $1 AND evm_chain_id = $2`, end, ubig.New(o.chainID))
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

func (o *orm) DeleteLogsAndBlocksAfter(ctx context.Context, start int64) error {
	// These deletes are bounded by reorg depth, so they are
	// fast and should not slow down the log readers.
	return o.Transaction(ctx, func(orm *orm) error {
		// Applying upper bound filter is critical for Postgres performance (especially for evm.logs table)
		// because it allows the planner to properly estimate the number of rows to be scanned.
		// If not applied, these queries can become very slow. After some critical number
		// of logs, Postgres will try to scan all the logs in the index by block_number.
		// Latency without upper bound filter can be orders of magnitude higher for large number of logs.
		_, err := o.db.ExecContext(ctx, `DELETE FROM evm.log_poller_blocks 
       						WHERE evm_chain_id = $1
       						AND block_number >= $2
       						AND block_number <= (SELECT MAX(block_number) 
						 		FROM evm.log_poller_blocks 
						 		WHERE evm_chain_id = $1)`,
			ubig.New(o.chainID), start)
		if err != nil {
			o.lggr.Warnw("Unable to clear reorged blocks, retrying", "err", err)
			return err
		}

		_, err = o.db.ExecContext(ctx, `DELETE FROM evm.logs 
       						WHERE evm_chain_id = $1 
       						AND block_number >= $2
       						AND block_number <= (SELECT MAX(block_number) FROM evm.logs WHERE evm_chain_id = $1)`,
			ubig.New(o.chainID), start)
		if err != nil {
			o.lggr.Warnw("Unable to clear reorged logs, retrying", "err", err)
			return err
		}
		return nil
	})
}

type Exp struct {
	Address      common.Address
	EventSig     common.Hash
	Expiration   time.Time
	TimeNow      time.Time
	ShouldDelete bool
}

func (o *orm) DeleteExpiredLogs(ctx context.Context, limit int64) (int64, error) {
	var err error
	var result sql.Result
	if limit > 0 {
		result, err = o.db.ExecContext(ctx, `
		DELETE FROM evm.logs
		WHERE (evm_chain_id, address, event_sig, block_number) IN (
			SELECT l.evm_chain_id, l.address, l.event_sig, l.block_number
			FROM evm.logs l
			INNER JOIN (
				SELECT address, event, MAX(retention) AS retention
				FROM evm.log_poller_filters
				WHERE evm_chain_id = $1
				GROUP BY evm_chain_id, address, event
				HAVING NOT 0 = ANY(ARRAY_AGG(retention))
			) r ON l.evm_chain_id = $1 AND l.address = r.address AND l.event_sig = r.event
			AND l.block_timestamp <= STATEMENT_TIMESTAMP() - (r.retention / 10^9 * interval '1 second')
			LIMIT $2
		)`, ubig.New(o.chainID), limit)
	} else {
		result, err = o.db.ExecContext(ctx, `WITH r AS
		( SELECT address, event, MAX(retention) AS retention
			FROM evm.log_poller_filters WHERE evm_chain_id=$1 
			GROUP BY evm_chain_id,address, event HAVING NOT 0 = ANY(ARRAY_AGG(retention))
		) DELETE FROM evm.logs l USING r
			WHERE l.evm_chain_id = $1 AND l.address=r.address AND l.event_sig=r.event
			AND l.created_at <= STATEMENT_TIMESTAMP() - (r.retention / 10^9 * interval '1 second')`, // retention is in nanoseconds (time.Duration aka BIGINT)
			ubig.New(o.chainID))
	}

	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

// InsertLogs is idempotent to support replays.
func (o *orm) InsertLogs(ctx context.Context, logs []Log) error {
	if err := o.validateLogs(logs); err != nil {
		return err
	}
	return o.Transaction(ctx, func(orm *orm) error {
		return o.insertLogsWithinTx(ctx, logs, orm.db.(*sqlx.Tx))
	})
}

func (o *orm) InsertLogsWithBlock(ctx context.Context, logs []Log, block LogPollerBlock) error {
	// Optimization, don't open TX when there is only a block to be persisted
	if len(logs) == 0 {
		return o.InsertBlock(ctx, block.BlockHash, block.BlockNumber, block.BlockTimestamp, block.FinalizedBlockNumber)
	}

	if err := o.validateLogs(logs); err != nil {
		return err
	}

	// Block and logs goes with the same TX to ensure atomicity
	return o.Transaction(ctx, func(orm *orm) error {
		if err := o.insertBlockWithinTx(ctx, orm.db.(*sqlx.Tx), block.BlockHash, block.BlockNumber, block.BlockTimestamp, block.FinalizedBlockNumber); err != nil {
			return err
		}
		return o.insertLogsWithinTx(ctx, logs, orm.db.(*sqlx.Tx))
	})
}

func (o *orm) insertBlockWithinTx(ctx context.Context, tx *sqlx.Tx, blockHash common.Hash, blockNumber int64, blockTimestamp time.Time, finalizedBlock int64) error {
	query := `INSERT INTO evm.log_poller_blocks (evm_chain_id, block_hash, block_number, block_timestamp, finalized_block_number, created_at)
			VALUES ($1, $2, $3, $4, $5, NOW())
			ON CONFLICT DO NOTHING`
	_, err := tx.ExecContext(ctx, query, ubig.New(o.chainID), blockHash.Bytes(), blockNumber, blockTimestamp, finalizedBlock)
	return err
}

func (o *orm) insertLogsWithinTx(ctx context.Context, logs []Log, tx *sqlx.Tx) error {
	batchInsertSize := 4000
	for i := 0; i < len(logs); i += batchInsertSize {
		start, end := i, i+batchInsertSize
		if end > len(logs) {
			end = len(logs)
		}

		_, err := tx.NamedExecContext(ctx, `
				INSERT INTO evm.logs 
					(evm_chain_id, log_index, block_hash, block_number, block_timestamp, address, event_sig, topics, tx_hash, data, created_at) 
				VALUES 
					(:evm_chain_id, :log_index, :block_hash, :block_number, :block_timestamp, :address, :event_sig, :topics, :tx_hash, :data, NOW()) 
				ON CONFLICT DO NOTHING`,
			logs[start:end],
		)

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) && batchInsertSize > 500 {
				// In case of DB timeouts, try to insert again with a smaller batch upto a limit
				batchInsertSize /= 2
				i -= batchInsertSize // counteract +=batchInsertSize on next loop iteration
				continue
			}
			return err
		}
	}
	return nil
}

func (o *orm) validateLogs(logs []Log) error {
	for _, log := range logs {
		if o.chainID.Cmp(log.EvmChainId.ToInt()) != 0 {
			return errors.Errorf("invalid chainID in log got %v want %v", log.EvmChainId.ToInt(), o.chainID)
		}
	}
	return nil
}

func (o *orm) SelectLogsByBlockRange(ctx context.Context, start, end int64) ([]Log, error) {
	args, err := newQueryArgs(o.chainID).
		withStartBlock(start).
		withEndBlock(end).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.logs 
        	WHERE evm_chain_id = :evm_chain_id
        	AND block_number >= :start_block 
        	AND block_number <= :end_block 
        	ORDER BY (block_number, log_index)`

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &logs, query, sqlArgs...)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogs finds the logs in a given block range.
func (o *orm) SelectLogs(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withStartBlock(start).
		withEndBlock(end).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id 
			AND address = :address
			AND event_sig = :event_sig  
			AND block_number >= :start_block 
			AND block_number <= :end_block
			ORDER BY (block_number, log_index)`

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &logs, query, sqlArgs...)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogsCreatedAfter finds logs created after some timestamp.
func (o *orm) SelectLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, after time.Time, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withBlockTimestampAfter(after).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
				WHERE evm_chain_id = :evm_chain_id
				AND address = :address
				AND event_sig = :event_sig
				AND block_timestamp > :block_timestamp_after
				AND block_number <= %s
				ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogsWithSigs finds the logs in the given block range with the given event signatures
// emitted from the given address.
func (o *orm) SelectLogsWithSigs(ctx context.Context, start, end int64, address common.Address, eventSigs []common.Hash) (logs []Log, err error) {
	args, err := newQueryArgs(o.chainID).
		withAddress(address).
		withEventSigArray(eventSigs).
		withStartBlock(start).
		withEndBlock(end).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.logs
				WHERE evm_chain_id = :evm_chain_id
				AND address = :address
				AND event_sig = ANY(:event_sig_array)
				AND block_number BETWEEN :start_block AND :end_block
				ORDER BY (block_number, log_index)`

	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &logs, query, sqlArgs...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return logs, err
}

func (o *orm) GetBlocksRange(ctx context.Context, start int64, end int64) ([]LogPollerBlock, error) {
	args, err := newQueryArgs(o.chainID).
		withStartBlock(start).
		withEndBlock(end).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.log_poller_blocks 
			WHERE block_number >= :start_block 
			AND block_number <= :end_block
			AND evm_chain_id = :evm_chain_id
			ORDER BY block_number ASC`

	var blocks []LogPollerBlock
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &blocks, query, sqlArgs...)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// SelectLatestLogEventSigsAddrsWithConfs finds the latest log by (address, event) combination that matches a list of Addresses and list of events
func (o *orm) SelectLatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, addresses []common.Address, eventSigs []common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgs(o.chainID).
		withAddressArray(addresses).
		withEventSigArray(eventSigs).
		withStartBlock(fromBlock).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs WHERE (block_number, address, event_sig) IN (
			SELECT MAX(block_number), address, event_sig FROM evm.logs 
				WHERE evm_chain_id = :evm_chain_id 
				AND event_sig = ANY(:event_sig_array) 
				AND address = ANY(:address_array) 
				AND block_number > :start_block 
				AND block_number <= %s
			GROUP BY event_sig, address
		)
		ORDER BY block_number ASC`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}
	return logs, nil
}

// SelectLatestBlockByEventSigsAddrsWithConfs finds the latest block number that matches a list of Addresses and list of events. It returns 0 if there is no matching block
func (o *orm) SelectLatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs Confirmations) (int64, error) {
	args, err := newQueryArgs(o.chainID).
		withEventSigArray(eventSigs).
		withAddressArray(addresses).
		withStartBlock(fromBlock).
		withConfs(confs).
		toArgs()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf(`
		SELECT COALESCE(MAX(block_number), 0) FROM evm.logs
			WHERE evm_chain_id = :evm_chain_id 
			AND event_sig = ANY(:event_sig_array) 
			AND address = ANY(:address_array) 
			AND block_number > :start_block 
			AND block_number <= %s`, nestedBlockNumberQuery(confs))

	var blockNumber int64
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.GetContext(ctx, &blockNumber, query, sqlArgs...); err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (o *orm) SelectLogsDataWordRange(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin, wordValueMax common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withWordIndex(wordIndex).
		withWordValueMin(wordValueMin).
		withWordValueMax(wordValueMax).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address 
			AND event_sig = :event_sig
			AND substring(data from 32*:word_index+1 for 32) >= :word_value_min
			AND substring(data from 32*:word_index+1 for 32) <= :word_value_max
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectLogsDataWordGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withWordIndex(wordIndex).
		withWordValueMin(wordValueMin).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address
			AND event_sig = :event_sig
			AND substring(data from 32*:word_index+1 for 32) >= :word_value_min
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectLogsDataWordBetween(ctx context.Context, address common.Address, eventSig common.Hash, wordIndexMin int, wordIndexMax int, wordValue common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withWordIndexMin(wordIndexMin).
		withWordIndexMax(wordIndexMax).
		withWordValue(wordValue).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address
			AND event_sig = :event_sig
			AND substring(data from 32*:word_index_min+1 for 32) <= :word_value
			AND substring(data from 32*:word_index_max+1 for 32) >= :word_value
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectIndexedLogsTopicGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withTopicIndex(topicIndex).
		withTopicValueMin(topicValueMin).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address 
			AND event_sig = :event_sig
			AND topics[:topic_index] >= :topic_value_min
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectIndexedLogsTopicRange(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin, topicValueMax common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withTopicIndex(topicIndex).
		withTopicValueMin(topicValueMin).
		withTopicValueMax(topicValueMax).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
			SELECT * FROM evm.logs 
				WHERE evm_chain_id = :evm_chain_id
				AND address = :address
				AND event_sig = :event_sig
				AND topics[:topic_index] >= :topic_value_min
				AND topics[:topic_index] <= :topic_value_max
				AND block_number <= %s
			ORDER BY (evm.logs.block_number, evm.logs.log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectIndexedLogs(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withTopicIndex(topicIndex).
		withTopicValues(topicValues).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address
			AND event_sig = :event_sig
			AND topics[:topic_index] = ANY(:topic_values)
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectIndexedLogsByBlockRange finds the indexed logs in a given block range.
func (o *orm) SelectIndexedLogsByBlockRange(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withTopicIndex(topicIndex).
		withTopicValues(topicValues).
		withStartBlock(start).
		withEndBlock(end).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.logs 
				WHERE evm_chain_id = :evm_chain_id 
				AND address = :address
				AND event_sig = :event_sig
				AND topics[:topic_index] = ANY(:topic_values)
				AND block_number >= :start_block
				AND block_number <= :end_block
				ORDER BY (block_number, log_index)`

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &logs, query, sqlArgs...)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectIndexedLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, after time.Time, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgsForEvent(o.chainID, address, eventSig).
		withBlockTimestampAfter(after).
		withConfs(confs).
		withTopicIndex(topicIndex).
		withTopicValues(topicValues).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address
			AND event_sig = :event_sig
			AND topics[:topic_index] = ANY(:topic_values)
			AND block_timestamp > :block_timestamp_after
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs))

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *orm) SelectIndexedLogsByTxHash(ctx context.Context, address common.Address, eventSig common.Hash, txHash common.Hash) ([]Log, error) {
	args, err := newQueryArgs(o.chainID).
		withTxHash(txHash).
		withAddress(address).
		withEventSig(eventSig).
		toArgs()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM evm.logs 
			WHERE evm_chain_id = :evm_chain_id
			AND address = :address
			AND event_sig = :event_sig
			AND tx_hash = :tx_hash
			ORDER BY (block_number, log_index)`

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	err = o.db.SelectContext(ctx, &logs, query, sqlArgs...)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectIndexedLogsWithSigsExcluding query's for logs that have signature A and exclude logs that have a corresponding signature B, matching is done based on the topic index both logs should be inside the block range and have the minimum number of confirmations
func (o *orm) SelectIndexedLogsWithSigsExcluding(ctx context.Context, sigA, sigB common.Hash, topicIndex int, address common.Address, startBlock, endBlock int64, confs Confirmations) ([]Log, error) {
	args, err := newQueryArgs(o.chainID).
		withAddress(address).
		withTopicIndex(topicIndex).
		withStartBlock(startBlock).
		withEndBlock(endBlock).
		withCustomHashArg("sigA", sigA).
		withCustomHashArg("sigB", sigB).
		withConfs(confs).
		toArgs()
	if err != nil {
		return nil, err
	}

	nestedQuery := nestedBlockNumberQuery(confs)
	query := fmt.Sprintf(`
		SELECT * FROM   evm.logs
		WHERE   evm_chain_id = :evm_chain_id
		AND     address = :address
		AND     event_sig = :sigA
		AND 	block_number BETWEEN :start_block AND :end_block
		AND 	block_number <= %s		
		EXCEPT
		SELECT     a.* FROM       evm.logs AS a
		INNER JOIN evm.logs B
		ON         a.evm_chain_id = b.evm_chain_id
		AND        a.address = b.address
		AND        a.topics[:topic_index] = b.topics[:topic_index]
		AND        a.event_sig = :sigA
		AND        b.event_sig = :sigB
	    AND 	   b.block_number BETWEEN :start_block AND :end_block
		AND		   b.block_number <= %s
		ORDER BY block_number,log_index ASC`, nestedQuery, nestedQuery)

	var logs []Log
	query, sqlArgs, _ := o.db.BindNamed(query, args)
	if err := o.db.SelectContext(ctx, &logs, query, sqlArgs...); err != nil {
		return nil, err
	}
	return logs, nil
}

func nestedBlockNumberQuery(confs Confirmations) string {
	if confs == Finalized {
		return `
				(SELECT finalized_block_number 
				FROM evm.log_poller_blocks 
				WHERE evm_chain_id = :evm_chain_id 
				ORDER BY block_number DESC LIMIT 1) `
	}
	// Intentionally wrap with greatest() function and don't return negative block numbers when :confs > :block_number
	// It doesn't impact logic of the outer query, because block numbers are never less or equal to 0 (guarded by log_poller_blocks_block_number_check)
	return `
			(SELECT greatest(block_number - :confs, 0) 
			FROM evm.log_poller_blocks 	
			WHERE evm_chain_id = :evm_chain_id 
			ORDER BY block_number DESC LIMIT 1) `
}

func UseTopicIndex(index int) (int, error) {
	// Only topicIndex 1 through 3 is valid. 0 is the event sig and only 4 total topics are allowed
	if !(index == 1 || index == 2 || index == 3) {
		return 0, fmt.Errorf("invalid index for topic: %d", index)
	}
	// Add 1 since postgresql arrays are 1-indexed.
	return index + 1, nil
}
package logpoller

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
)

// TODO: Set a reasonable timeout
const defaultTimeout = 10 * time.Second

// ORM represents the persistent data access layer used by the log poller. At this moment, it's a bit leaky abstraction, because
// it exposes some of the database implementation details (e.g. pg.Q). Ideally it should be agnostic and could be applied to any persistence layer.
// What is more, LogPoller should not be aware of the underlying database implementation and delegate all the queries to the ORM.
type ORM interface {
	InsertLogs(ctx context.Context, logs []Log) error
	InsertLogsWithBlock(ctx context.Context, logs []Log, block LogPollerBlock) error
	InsertFilter(ctx context.Context, filter Filter) error

	LoadFilters(ctx context.Context) (map[string]Filter, error)
	DeleteFilter(ctx context.Context, name string) error

	DeleteBlocksBefore(ctx context.Context, end int64) error
	DeleteLogsAndBlocksAfter(ctx context.Context, start int64) error
	DeleteExpiredLogs(ctx context.Context) error

	GetBlocksRange(ctx context.Context, start int64, end int64) ([]LogPollerBlock, error)
	SelectBlockByNumber(ctx context.Context, blockNumber int64) (*LogPollerBlock, error)
	SelectLatestBlock(ctx context.Context) (*LogPollerBlock, error)

	SelectLogs(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash) ([]Log, error)
	SelectLogsWithSigs(ctx context.Context, start, end int64, address common.Address, eventSigs []common.Hash) ([]Log, error)
	SelectLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, after time.Time, confs Confirmations) ([]Log, error)
	SelectLatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs Confirmations) (*Log, error)
	SelectLatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, addresses []common.Address, eventSigs []common.Hash, confs Confirmations) ([]Log, error)
	SelectLatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs Confirmations) (int64, error)

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

type DbORM struct {
	chainID *big.Int
	db      sqlutil.Queryer
	lggr    logger.Logger
}

var _ ORM = &DbORM{}

// NewORM creates a DbORM scoped to chainID.
func NewORM(chainID *big.Int, db *sqlx.DB, lggr logger.Logger) *DbORM {
	return &DbORM{
		chainID: chainID,
		db:      db,
		lggr:    lggr,
	}
}

// InsertBlock is idempotent to support replays.
func (o *DbORM) InsertBlock(ctx context.Context, blockHash common.Hash, blockNumber int64, blockTimestamp time.Time, finalizedBlock int64) error {
	query := `INSERT INTO evm.log_poller_blocks (evm_chain_id, block_hash, block_number, block_timestamp, finalized_block_number, created_at)
			VALUES ($1, $2, $3, $4, $5, NOW())
			ON CONFLICT DO NOTHING`
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// TODO: Why do we get a transaction failure?? Not using a transaction..
	// TODO: Maybe because the previous transaction failed for SelectLatestLogByEventSigWithConfs
	_, err := o.db.ExecContext(ctx, query, o.chainID.String(), blockHash, blockNumber, blockTimestamp, finalizedBlock)
	return err
}

// InsertFilter is idempotent.
//
// Each address/event pair must have a unique job id, so it may be removed when the job is deleted.
// If a second job tries to overwrite the same pair, this should fail.
func (o *DbORM) InsertFilter(ctx context.Context, filter Filter) (err error) {
	// '::' has to be escaped in the query string
	// https://github.com/jmoiron/sqlx/issues/91, https://github.com/jmoiron/sqlx/issues/428
	query := `
		INSERT INTO evm.log_poller_filters
	  		(name, evm_chain_id, retention, created_at, address, event)
		SELECT * FROM
			(SELECT $1, $2 ::NUMERIC, $3 ::BIGINT, NOW()) x,
			(SELECT unnest($4 ::BYTEA[]) addr) a,
			(SELECT unnest($5 ::BYTEA[]) ev) e
		ON CONFLICT (name, evm_chain_id, address, event) 
		DO UPDATE SET retention=$3 ::BIGINT`

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err = o.db.ExecContext(ctx, query, filter.Name, o.chainID.String(), filter.Retention, concatBytes(filter.Addresses), concatBytes(filter.EventSigs))
	return err
}

// DeleteFilter removes all events,address pairs associated with the Filter
func (o *DbORM) DeleteFilter(ctx context.Context, name string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := o.db.ExecContext(ctx,
		`DELETE FROM evm.log_poller_filters WHERE name = $1 AND evm_chain_id = $2`,
		name, ubig.New(o.chainID))
	return err

}

// LoadFilters returns all filters for this chain
func (o *DbORM) LoadFilters(ctx context.Context) (map[string]Filter, error) {
	rows := make([]Filter, 0)

	query := `SELECT name,
			ARRAY_AGG(DISTINCT address)::BYTEA[] AS addresses, 
			ARRAY_AGG(DISTINCT event)::BYTEA[] AS event_sigs,
			MAX(retention) AS retention
		FROM evm.log_poller_filters WHERE evm_chain_id = $1
		GROUP BY name`

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &rows, query, ubig.New(o.chainID))
	/*
		err := q.Select(&rows, `SELECT name,
				ARRAY_AGG(DISTINCT address)::BYTEA[] AS addresses,
				ARRAY_AGG(DISTINCT event)::BYTEA[] AS event_sigs,
				MAX(retention) AS retention
			FROM evm.log_poller_filters WHERE evm_chain_id = $1
			GROUP BY name`, ubig.New(o.chainID))
	*/
	filters := make(map[string]Filter)
	for _, filter := range rows {
		filters[filter.Name] = filter
	}

	return filters, err
}

func (o *DbORM) SelectBlockByHash(ctx context.Context, hash common.Hash) (*LogPollerBlock, error) {
	var b LogPollerBlock
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE block_hash = $1 AND evm_chain_id = $2`, hash, ubig.New(o.chainID)); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *DbORM) SelectBlockByNumber(ctx context.Context, n int64) (*LogPollerBlock, error) {
	var b LogPollerBlock
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE block_number = $1 AND evm_chain_id = $2`, n, ubig.New(o.chainID)); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *DbORM) SelectLatestBlock(ctx context.Context) (*LogPollerBlock, error) {
	var b LogPollerBlock
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.GetContext(ctx, &b, `SELECT * FROM evm.log_poller_blocks WHERE evm_chain_id = $1 ORDER BY block_number DESC LIMIT 1`, o.chainID.String()); err != nil {
		return nil, err
	}
	return &b, nil
}

func (o *DbORM) SelectLatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs Confirmations) (*Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs
			WHERE evm_chain_id = $1
			AND event_sig = $2
			AND address = $3
			AND block_number <= %s
			ORDER BY (block_number, log_index) DESC LIMIT 1`,
		nestedBlockNumberQuery(confs, o.chainID.String()))
	var l Log

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.GetContext(ctx, &l, query, o.chainID.String(), eventSig, address); err != nil {
		return nil, err
	}
	return &l, nil
}

// DeleteBlocksBefore delete all blocks before and including end.
func (o *DbORM) DeleteBlocksBefore(ctx context.Context, end int64) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err := o.db.ExecContext(ctx, `DELETE FROM evm.log_poller_blocks WHERE block_number <= $1 AND evm_chain_id = $2`, end, ubig.New(o.chainID))
	return err
}

func (o *DbORM) DeleteLogsAndBlocksAfter(ctx context.Context, start int64) error {
	// These deletes are bounded by reorg depth, so they are
	// fast and should not slow down the log readers.
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	// TODO: Is Transact working?? Why are tests failing
	performInsert := func(tx *sqlx.Tx) error {
		args, err := newQueryArgs(o.chainID).
			withStartBlock(start).
			toArgs()
		if err != nil {
			o.lggr.Error("Cant build args for DeleteLogsAndBlocksAfter queries", "err", err)
			return err
		}

		_, err = tx.NamedExec(`DELETE FROM evm.log_poller_blocks WHERE block_number >= :start_block AND evm_chain_id = :evm_chain_id`, args)
		if err != nil {
			o.lggr.Warnw("Unable to clear reorged blocks, retrying", "err", err)
			return err
		}

		_, err = tx.NamedExec(`DELETE FROM evm.logs WHERE block_number >= :start_block AND evm_chain_id = :evm_chain_id`, args)
		if err != nil {
			o.lggr.Warnw("Unable to clear reorged logs, retrying", "err", err)
			return err
		}
		return nil
	}
	return sqlutil.Transact[*sqlx.Tx](ctx, func(q sqlutil.Queryer) *sqlx.Tx {
		return q.(*sqlx.Tx)
	}, o.db, nil, performInsert)
}

type Exp struct {
	Address      common.Address
	EventSig     common.Hash
	Expiration   time.Time
	TimeNow      time.Time
	ShouldDelete bool
}

func (o *DbORM) DeleteExpiredLogs(ctx context.Context) error {
	// TODO: LongQueryTimeout?
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err := o.db.ExecContext(ctx, `WITH r AS
		( SELECT address, event, MAX(retention) AS retention
			FROM evm.log_poller_filters WHERE evm_chain_id=$1 
			GROUP BY evm_chain_id,address, event HAVING NOT 0 = ANY(ARRAY_AGG(retention))
		) DELETE FROM evm.logs l USING r
			WHERE l.evm_chain_id = $1 AND l.address=r.address AND l.event_sig=r.event
			AND l.created_at <= STATEMENT_TIMESTAMP() - (r.retention / 10^9 * interval '1 second')`, // retention is in nanoseconds (time.Duration aka BIGINT)
		ubig.New(o.chainID))
	return err
}

// InsertLogs is idempotent to support replays.
func (o *DbORM) InsertLogs(ctx context.Context, logs []Log) error {
	if err := o.validateLogs(logs); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	performInsert := func(tx *sqlx.Tx) error {
		return o.insertLogsWithinTx(ctx, logs, tx)
	}
	return sqlutil.Transact[*sqlx.Tx](ctx, func(q sqlutil.Queryer) *sqlx.Tx {
		return q.(*sqlx.Tx)
	}, o.db, nil, performInsert)
}

func (o *DbORM) InsertLogsWithBlock(ctx context.Context, logs []Log, block LogPollerBlock) error {
	// Optimization, don't open TX when there is only a block to be persisted
	if len(logs) == 0 {
		return o.InsertBlock(ctx, block.BlockHash, block.BlockNumber, block.BlockTimestamp, block.FinalizedBlockNumber)
	}

	if err := o.validateLogs(logs); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// Block and logs goes with the same TX to ensure atomicity
	performInsert := func(tx *sqlx.Tx) error {
		if err := o.InsertBlock(ctx, block.BlockHash, block.BlockNumber, block.BlockTimestamp, block.FinalizedBlockNumber); err != nil {
			return err
		}
		return o.insertLogsWithinTx(ctx, logs, tx)
	}
	return sqlutil.Transact[*sqlx.Tx](ctx, func(q sqlutil.Queryer) *sqlx.Tx {
		return q.(*sqlx.Tx)
	}, o.db, nil, performInsert)
}

func (o *DbORM) insertLogsWithinTx(ctx context.Context, logs []Log, tx *sqlx.Tx) error {
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

func (o *DbORM) validateLogs(logs []Log) error {
	for _, log := range logs {
		if o.chainID.Cmp(log.EvmChainId.ToInt()) != 0 {
			return errors.Errorf("invalid chainID in log got %v want %v", log.EvmChainId.ToInt(), o.chainID)
		}
	}
	return nil
}

func (o *DbORM) SelectLogsByBlockRange(ctx context.Context, start, end int64) ([]Log, error) {
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &logs, `
        SELECT * FROM evm.logs 
        	WHERE evm_chain_id = $1
        	AND block_number >= $2 
        	AND block_number <= $3 
        	ORDER BY (block_number, log_index, created_at)`, o.chainID.String(), start, end)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogs finds the logs in a given block range.
func (o *DbORM) SelectLogs(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash) ([]Log, error) {
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &logs, `
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1 
			AND address = $2
			AND event_sig = $3  
			AND block_number >= $4 
			AND block_number <= $5
			ORDER BY (block_number, log_index)`, o.chainID.String(), address, eventSig, start, end)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogsCreatedAfter finds logs created after some timestamp.
func (o *DbORM) SelectLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, after time.Time, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
				WHERE evm_chain_id = $1
				AND address = $2
				AND event_sig = $3
				AND block_timestamp > $4
				AND block_number <= %s
				ORDER BY (block_number, log_index)`,
		nestedBlockNumberQuery(confs, o.chainID.String()))

	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, after); err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectLogsWithSigs finds the logs in the given block range with the given event signatures
// emitted from the given address.
func (o *DbORM) SelectLogsWithSigs(ctx context.Context, start, end int64, address common.Address, eventSigs []common.Hash) (logs []Log, err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err = o.db.SelectContext(ctx, &logs, `
			SELECT * FROM evm.logs
				WHERE evm_chain_id = $1
				AND address = $2
				AND event_sig = ANY($3)
				AND block_number BETWEEN $4 AND $5
				ORDER BY (block_number, log_index)`, o.chainID.String(), address, concatBytes(eventSigs), start, end)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return logs, err
}

func (o *DbORM) GetBlocksRange(ctx context.Context, start int64, end int64) ([]LogPollerBlock, error) {
	var blocks []LogPollerBlock
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &blocks, `
        SELECT * FROM evm.log_poller_blocks 
			WHERE block_number >= $1 
			AND block_number <= $2
			AND evm_chain_id = $3
			ORDER BY block_number ASC`, start, end, o.chainID.String())
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

// SelectLatestLogEventSigsAddrsWithConfs finds the latest log by (address, event) combination that matches a list of Addresses and list of events
func (o *DbORM) SelectLatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, addresses []common.Address, eventSigs []common.Hash, confs Confirmations) ([]Log, error) {
	// TODO: cant convert byteArray!?
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs WHERE (block_number, address, event_sig) IN (
			SELECT MAX(block_number), address, event_sig FROM evm.logs 
				WHERE evm_chain_id = $1 
				AND event_sig = ANY($2) 
				AND address = ANY($3) 
				AND block_number > $4
				AND block_number <= %s
			GROUP BY event_sig, address
		)
		ORDER BY block_number ASC`, nestedBlockNumberQuery(confs, o.chainID.String()))

	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), concatBytes(eventSigs), concatBytes(addresses), fromBlock); err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}
	return logs, nil
}

// SelectLatestBlockByEventSigsAddrsWithConfs finds the latest block number that matches a list of Addresses and list of events. It returns 0 if there is no matching block
func (o *DbORM) SelectLatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs Confirmations) (int64, error) {
	query := fmt.Sprintf(`
		SELECT COALESCE(MAX(block_number), 0) FROM evm.logs
			WHERE evm_chain_id = $1 
			AND event_sig = ANY($2) 
			AND address = ANY($3) 
			AND block_number > $4 
			AND block_number <= %s`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var blockNumber int64
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.GetContext(ctx, &blockNumber, query, o.chainID.String(), concatBytes(eventSigs), concatBytes(addresses), fromBlock); err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (o *DbORM) SelectLogsDataWordRange(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin, wordValueMax common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2 
			AND event_sig = $3
			AND substring(data from 32*$4+1 for 32) >= $5
			AND substring(data from 32*$4+1 for 32) <= $6
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, wordIndex, wordValueMin, wordValueMax); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectLogsDataWordGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2
			AND event_sig = $3
			AND substring(data from 32*$4+1 for 32) >= $5
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, wordIndex, wordValueMin); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectLogsDataWordBetween(ctx context.Context, address common.Address, eventSig common.Hash, wordIndexMin int, wordIndexMax int, wordValue common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2
			AND event_sig = $3
			AND substring(data from 32*$4+1 for 32) <= $5
			AND substring(data from 32*$6+1 for 32) >= $5
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, wordIndexMin, wordValue, wordIndexMax); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectIndexedLogsTopicGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs
			WHERE evm_chain_id = $1
			AND address = $2 
			AND event_sig = $3
			AND topics[$4] >= $5
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, topicIndex, topicValueMin); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectIndexedLogsTopicRange(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin, topicValueMax common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
			SELECT * FROM evm.logs 
				WHERE evm_chain_id = $1
				AND address = $2
				AND event_sig = $3
				AND topics[$4] >= $5
				AND topics[$4] <= $6
				AND block_number <= %s
			ORDER BY (evm.logs.block_number, evm.logs.log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, topicIndex, topicValueMin, topicValueMax); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectIndexedLogs(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2
			AND event_sig = $3
			AND topics[$4] = ANY($5)
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, topicIndex, concatBytes(topicValues)); err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectIndexedLogsByBlockRange finds the indexed logs in a given block range.
func (o *DbORM) SelectIndexedLogsByBlockRange(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash) ([]Log, error) {
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &logs, `
		SELECT * FROM evm.logs 
				WHERE evm_chain_id = $1 
				AND address = $2
				AND event_sig = $3
				AND topics[$4] = ANY($5)
				AND block_number >= $6
				AND block_number <= $7
				ORDER BY (block_number, log_index)`,
		o.chainID.String(), address, eventSig, topicIndex, concatBytes(topicValues), start, end)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectIndexedLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, after time.Time, confs Confirmations) ([]Log, error) {
	query := fmt.Sprintf(`
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2
			AND event_sig = $3
			AND topics[$4] = ANY($5)
			AND block_timestamp > $6
			AND block_number <= %s
			ORDER BY (block_number, log_index)`, nestedBlockNumberQuery(confs, o.chainID.String()))

	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, eventSig, topicIndex, concatBytes(topicValues), after); err != nil {
		return nil, err
	}
	return logs, nil
}

func (o *DbORM) SelectIndexedLogsByTxHash(ctx context.Context, address common.Address, eventSig common.Hash, txHash common.Hash) ([]Log, error) {
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	err := o.db.SelectContext(ctx, &logs, `
		SELECT * FROM evm.logs 
			WHERE evm_chain_id = $1
			AND address = $2
			AND event_sig = $3			  
			AND tx_hash = $4
			ORDER BY (block_number, log_index)`, o.chainID.String(), address, eventSig, txHash)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// SelectIndexedLogsWithSigsExcluding query's for logs that have signature A and exclude logs that have a corresponding signature B, matching is done based on the topic index both logs should be inside the block range and have the minimum number of confirmations
func (o *DbORM) SelectIndexedLogsWithSigsExcluding(ctx context.Context, sigA, sigB common.Hash, topicIndex int, address common.Address, startBlock, endBlock int64, confs Confirmations) ([]Log, error) {
	nestedQuery := nestedBlockNumberQuery(confs, o.chainID.String())
	query := fmt.Sprintf(`
		SELECT * FROM   evm.logs
		WHERE   evm_chain_id = $1
		AND     address = $2
		AND     event_sig = $3
		AND 	block_number BETWEEN $5 AND $6
		AND 	block_number <= %s		
		EXCEPT
		SELECT     a.* FROM       evm.logs AS a
		INNER JOIN evm.logs B
		ON         a.evm_chain_id = b.evm_chain_id
		AND        a.address = b.address
		AND        a.topics[$7] = b.topics[$7]
		AND        a.event_sig = $3
		AND        b.event_sig = $4
	    AND 	   b.block_number BETWEEN $5 AND $6
		AND		   b.block_number <= %s
		ORDER BY block_number,log_index ASC`, nestedQuery, nestedQuery)
	var logs []Log
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	if err := o.db.SelectContext(ctx, &logs, query, o.chainID.String(), address, sigA, sigB, startBlock, endBlock, topicIndex); err != nil {
		return nil, err
	}
	return logs, nil
}

func nestedBlockNumberQuery(confs Confirmations, chainID string) string {
	if confs == Finalized {
		return fmt.Sprintf(`
				(SELECT finalized_block_number 
				FROM evm.log_poller_blocks 
				WHERE evm_chain_id = %s 
				ORDER BY block_number DESC LIMIT 1) `, chainID)
	}
	// Intentionally wrap with greatest() function and don't return negative block numbers when :confs > :block_number
	// It doesn't impact logic of the outer query, because block numbers are never less or equal to 0 (guarded by log_poller_blocks_block_number_check)
	return fmt.Sprintf(`
			(SELECT greatest(block_number - %d, 0) 
			FROM evm.log_poller_blocks 	
			WHERE evm_chain_id = %s 
			ORDER BY block_number DESC LIMIT 1) `, confs, chainID)

}

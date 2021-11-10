package pgtest

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/scylladb/go-reflectx"
	"github.com/smartcontractkit/go-txdb"
	"github.com/smartcontractkit/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/smartcontractkit/chainlink/core/logger"
)

func init() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("You must provide a DATABASE_URL environment variable")
	}

	parsed, err := url.Parse(dbURL)
	if err != nil {
		panic(err)
	}
	if parsed.Path == "" {
		msg := fmt.Sprintf("invalid DATABASE_URL: `%s`. You must set DATABASE_URL env var to point to your test database. Note that the test database MUST end in `_test` to differentiate from a possible production DB. HINT: Try DATABASE_URL=postgresql://postgres@localhost:5432/chainlink_test?sslmode=disable", parsed.String())
		panic(msg)
	}
	if !strings.HasSuffix(parsed.Path, "_test") {
		msg := fmt.Sprintf("cannot run tests against database named `%s`. Note that the test database MUST end in `_test` to differentiate from a possible production DB. HINT: Try DATABASE_URL=postgresql://postgres@localhost:5432/chainlink_test?sslmode=disable", parsed.Path[1:])
		panic(msg)
	}

	// Disable SavePoints because they cause random errors for reasons I cannot fathom.
	// NOTE: That this will cause transaction BEGIN/ROLLBACK to effectively be
	// a no-op, this should have no negative impact on normal test operation.
	// If you MUST test BEGIN/ROLLBACK behaviour, you will have to configure your
	// store to use the raw DialectPostgres dialect and setup a one-use database.
	// See BootstrapThrowawayORM() as a convenience function to help you do this.
	// TODO: re-enable savepoint emulation once gorm is removed:
	// https://app.clubhouse.io/chainlinklabs/story/8781/remove-dependency-on-gorm
	txdb.Register("txdb", "pgx", dbURL, txdb.SavePointOption(nil))
	sqlx.BindDriver("txdb", sqlx.DOLLAR)
}

func NewGormDB(t *testing.T) *gorm.DB {
	sqlDB := NewSqlDB(t)
	return GormDBFromSql(t, sqlDB)
}

func GormDBFromSql(t *testing.T, db *sql.DB) *gorm.DB {
	logAllQueries := os.Getenv("LOG_SQL") == "true"
	newLogger := logger.NewGormWrapper(logger.TestLogger(t), logAllQueries, 0)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
		DSN:  uuid.NewV4().String(),
	}), &gorm.Config{Logger: newLogger})

	require.NoError(t, err)

	// Incantation to fix https://github.com/go-gorm/gorm/issues/4586
	gormDB = gormDB.Omit(clause.Associations).Session(&gorm.Session{})

	return gormDB
}

func NewSqlDB(t *testing.T) *sql.DB {
	db, err := sql.Open("txdb", uuid.NewV4().String())
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, db.Close()) })

	// There is a bug to do with context cancellation somewhere in txdb or sql.
	// If you try to use the DB "too quickly" using a .WithContext then cancel
	// the context, the transaction state gets poisoned or lost somehow and
	// subsequent queries fail with "sql: transaction has already been
	// committed or rolled back" (although postgres does not log any errors).

	// Calling SELECT 1 here seems to reliably fix it. Created an issue to track here:
	// https://github.com/DATA-DOG/go-txdb/issues/43
	_, err = db.Exec(`SELECT 1`)
	require.NoError(t, err)

	return db
}

func NewSqlxDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("txdb", uuid.NewV4().String())
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, db.Close()) })

	// There is a bug to do with context cancellation somewhere in txdb or sql.
	// If you try to use the DB "too quickly" using a .WithContext then cancel
	// the context, the transaction state gets poisoned or lost somehow and
	// subsequent queries fail with "sql: transaction has already been
	// committed or rolled back" (although postgres does not log any errors).

	// Calling SELECT 1 here seems to reliably fix it. Created an issue to track here:
	// https://github.com/DATA-DOG/go-txdb/issues/43
	_, err = db.Exec(`SELECT 1`)
	require.NoError(t, err)

	db.MapperFunc(reflectx.CamelToSnakeASCII)

	return db
}

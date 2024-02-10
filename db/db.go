package db

import (
	"context"
	"database/sql"
	"errors"
	"restexample/logadapter"

	sqldblogger "github.com/simukti/sqldb-logger"

	_ "github.com/go-sql-driver/mysql"
)

type QueryAble interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

var db *sql.DB

func Open(dsn string, debug bool, mock bool) (err error) {

	if mock {
		return
	}

	if db, err = sql.Open("mysql", dsn); err != nil {
		return
	}

	if debug {
		loggerAdapter := logadapter.NewSimpleAdapter(&logadapter.SimpleLoggerAdapter{
			Logger: &logadapter.Logger{},
		})
		db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter,
			sqldblogger.WithSQLQueryAsMessage(true), // default: false
		) // db is STILL *sql.DB
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}

func DB() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}

func Begin() (*sql.Tx, error) {
	return db.Begin()
}

func Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func Commit(tx *sql.Tx) error {
	return tx.Commit()
}

var (
	ErrNoConnection = errors.New("expected a db connection or a transaction. Nil received")
)

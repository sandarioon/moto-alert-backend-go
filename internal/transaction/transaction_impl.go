package transaction

import (
	"context"
	"database/sql"
	"time"

	postgres "github.com/sandarioon/moto-alert-backend-go/pkg/database"
)

type SQLTransactioner struct {
	db *postgres.DBLogger
}

func NewSQLTransactioner(db *postgres.DBLogger) *SQLTransactioner {
	return &SQLTransactioner{db: db}
}

func (st *SQLTransactioner) BeginTx(ctx context.Context) (Transaction, error) {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &SQLTransaction{tx: tx}, nil
}

type SQLTransaction struct {
	tx *sql.Tx
}

func (st *SQLTransaction) Commit() error {
	start := time.Now()
	postgres.PrintSql(true, "TX. COMMIT;", &start)
	return st.tx.Commit()
}

func (st *SQLTransaction) Rollback() error {
	start := time.Now()
	postgres.PrintSql(false, "TX. ROLLBACK;", &start)
	return st.tx.Rollback()
}

func (st *SQLTransaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()

	result, err := st.tx.Exec(query, args...)
	if err != nil {
		postgres.PrintSql(false, "TX. "+query, &start, args...)
		return nil, err
	}

	return result, err
}

func (st *SQLTransaction) QueryRow(query string, args ...interface{}) *sql.Row {
	start := time.Now()

	row := st.tx.QueryRow(query, args...)
	postgres.PrintSql(true, "TX. "+query, &start, args...)
	return row
}
func (st *SQLTransaction) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	start := time.Now()

	row := st.tx.QueryRowContext(ctx, query, args...)
	postgres.PrintSql(true, "TX. "+query, &start, args...)
	return row
}

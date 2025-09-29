package transaction

import (
	"context"
	"database/sql"
)

type Transactioner interface {
	BeginTx(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	Commit() error
	Rollback() error
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

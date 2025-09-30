package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// DBLogger wraps around sql.DB to log queries
type DBLogger struct {
	*sql.DB
}

// NewDBLogger creates a new instance of DBLogger
func NewDBLogger(db *sql.DB) *DBLogger {
	return &DBLogger{db}
}

func (d *DBLogger) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	start := time.Now()
	tx, err := d.DB.BeginTx(ctx, opts)
	if err != nil {
		PrintSql(false, "TX. BEGIN;", &start)
		return nil, err
	}
	PrintSql(true, "TX. BEGIN;", &start)
	return tx, err
}

// Query overrides the default Query method to log the query
func (d *DBLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := d.DB.Query(query, args...)
	if err != nil {
		PrintSql(false, query, &start, args...)
		return nil, err
	}
	PrintSql(true, query, &start, args...)
	return rows, err
}

// QueryRow overrides the default QueryRow method to log the query
func (d *DBLogger) QueryRow(query string, args ...interface{}) *sql.Row {
	start := time.Now()
	row := d.DB.QueryRow(query, args...)
	PrintSql(true, query, &start, args...)
	return row
}

func (d *DBLogger) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	start := time.Now()
	row := d.DB.QueryRowContext(ctx, query, args...)
	PrintSql(true, query, &start, args...)
	return row
}

// Exec overrides the default Exec method to log the query
func (d *DBLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := d.DB.Exec(query, args...)
	if err != nil {
		PrintSql(false, query, &start, args...)
		return nil, err
	}
	PrintSql(true, query, &start, args...)
	return result, err
}

func (d *DBLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		PrintSql(false, query, &start, args...)
		return nil, err
	}
	PrintSql(true, query, &start, args...)
	return result, err
}

func NewDB(cfg Config) (*DBLogger, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Errorf("Unable to connect to database: %v", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf("Ping failed: %v", err)
		panic(err)
	}

	dbLogger := NewDBLogger(db)

	return dbLogger, nil
}

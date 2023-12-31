/*
	here, we define the interface, we can say also, we define the behavior of interface
	that will be implemented at postgres.go
*/

package database

import (
	"context"

	"github.com/jackc/pgconn"
)

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Transaction interface {
	Rollback(context.Context) error
	BulkInsert(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error)
	Commit(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
}

type DB interface {
	Close()
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Begin(ctx context.Context) (Transaction, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
}

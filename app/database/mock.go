// this for testing purpose
package database

import (
	"context"

	"github.com/jackc/pgconn"
)

type mockRow struct {
	MockScan func(dest ...any) error
}

func (this *mockRow) Scan(dest ...any) error {
	return this.MockScan(dest...)
}

func NewMockRow() mockRow {
	return mockRow{MockScan: func(dest ...any) error {
		return nil
	}}
}

type mockRows struct {
	MockClose func()
	MockErr   func() error
	MockNext  func() bool
	MockScan  func(dest ...any) error
}

func (this *mockRows) Close() {
	this.MockClose()
}

func (this *mockRows) Err() error {
	return this.MockErr()
}

func (this *mockRows) Next() bool {
	return this.MockNext()
}

func (this *mockRows) Scan(dest ...any) error {
	return this.MockScan(dest...)
}

func NewMockRows() mockRows {
	return mockRows{
		MockClose: func() {},
		MockErr:   func() error { return nil },
		MockNext:  func() bool { return true },
		MockScan:  func(dest ...any) error { return nil },
	}
}

type mockTransaction struct {
	MockRow        mockRow
	MockRows       mockRows
	MockRollback   func(ctx context.Context) error
	MockBulkInsert func(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error)
	MockCommit     func(ctx context.Context) error
	MockQueryRow   func(ctx context.Context, sql string, args ...any) Row
	MockQuery      func(ctx context.Context, sql string, args ...any) (Rows, error)
	MockExec       func(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
}

func (this *mockTransaction) Rollback(ctx context.Context) error {
	return this.MockRollback(ctx)
}

func (this *mockTransaction) BulkInsert(ctx context.Context, tableName string, columsn []string,
	rows [][]any) (int, error) {
	return this.MockBulkInsert(ctx, tableName, columsn, rows)
}

func (this *mockTransaction) Commit(ctx context.Context) error {
	return this.MockCommit(ctx)
}

func (this *mockTransaction) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return this.MockQueryRow(ctx, sql, args...)
}

func (this *mockTransaction) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return this.Query(ctx, sql, args...)
}

func (this *mockTransaction) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return this.Exec(ctx, sql, arguments...)
}

func NewMockTransaction() mockTransaction {
	mockRow := NewMockRow()
	mockRows := NewMockRows()
	return mockTransaction{
		MockRow:  mockRow,
		MockRows: mockRows,
		MockRollback: func(ctx context.Context) error {
			return nil
		},
		MockBulkInsert: func(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error) {
			return 0, nil
		},
		MockCommit: func(ctx context.Context) error {
			return nil
		},
		MockQuery: func(ctx context.Context, sql string, args ...any) (Rows, error) {
			return &mockRows, nil
		},
		MockQueryRow: func(ctx context.Context, sql string, args ...any) Row {
			return &mockRow
		},
		MockExec: func(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
			return pgconn.CommandTag{}, nil
		},
	}
}

type mockDB struct {
	MockRow         mockRow
	MockTransaction mockTransaction
	MockRows        mockRows
	MockClose       func()
	MockQueryRow    func(ctx context.Context, sql string, args ...any) Row
	MockBegin       func(ctx context.Context) (Transaction, error)
	MockQuery       func(ctx context.Context, sql string, args ...any) (Rows, error)
}

func (this *mockDB) Close() {
	this.MockClose()
}

func (this *mockDB) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return this.MockQueryRow(ctx, sql, args...)
}

func (this *mockDB) Begin(ctx context.Context) (Transaction, error) {
	return this.MockBegin(ctx)
}

func (this *mockDB) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return this.Query(ctx, sql, args...)
}

func NewMockDB() mockDB {
	mockRow := NewMockRow()
	mockTransaction := NewMockTransaction()
	mockRows := NewMockRows()
	return mockDB{
		MockRow:         mockRow,
		MockTransaction: mockTransaction,
		MockRows:        mockRows,
		MockClose:       func() {},
		MockQueryRow: func(ctx context.Context, sql string, args ...any) Row {
			return &mockRow
		},
		MockBegin: func(ctx context.Context) (Transaction, error) {
			return &mockTransaction, nil
		},
		MockQuery: func(ctx context.Context, sql string, args ...any) (Rows, error) {
			return &mockRows, nil
		},
	}
}

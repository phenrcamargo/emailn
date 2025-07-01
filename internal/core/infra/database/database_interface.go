package database

import "database/sql"

type DatabaseInterface interface {
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(tx *sql.Tx, query string, args ...any) (sql.Result, error)
}

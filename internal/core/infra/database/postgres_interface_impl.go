package database

import "database/sql"

type PostgresInterfaceImpl struct{}

func (p PostgresInterfaceImpl) QueryRow(query string, args ...any) *sql.Row {
	db := OpenConnection()
	defer db.Close()

	return db.QueryRow(query, args...)
}

func (p PostgresInterfaceImpl) Query(query string, args ...any) (*sql.Rows, error) {
	db := OpenConnection()
	defer db.Close()

	return db.Query(query, args...)
}

func (p PostgresInterfaceImpl) Exec(tx *sql.Tx, query string, args ...any) (sql.Result, error) {
	return tx.Exec(query, args...)
}

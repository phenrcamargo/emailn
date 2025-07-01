package database

import "database/sql"

func OpenConnection() *sql.DB {
	connStr := "user=postgres dbname=emailn sslmode=disable password=postgres host=localhost port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Fail to connect database -> " + err.Error())
	}

	return db
}

func Begin() (*sql.Tx, error) {
	db := OpenConnection()
	return db.Begin()
}

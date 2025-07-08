package sqldb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteAdapter struct{}

func (a *SQLiteAdapter) Connect(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite3", dsn)	
}
package sqldb

import (
	"database/sql"
	"log"
)

var db *sql.DB

type DBAdapter interface {
	Connect(dsn string) (*sql.DB, error)
}

func Init(adapter DBAdapter, dsn string) (*sql.DB, error) {
	var err error
	db, err = adapter.Connect(dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected.")

	err = RunMigrations()
	if err != nil {
		log.Printf("Failed to run migrations: %v\n", err)
		return nil, err
	}

	return db, nil
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close DB: %v\n", err)
		} else {
			log.Println("Database closed.")
		}
	}
}
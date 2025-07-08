package sqldb

import "log"

func RunMigrations() error {
	// This function is a placeholder for running database migrations.
	// In a real application, you would use a migration tool or library
	// to apply schema changes to your database.
	//
	// For example, you might use:
	// - golang-migrate/migrate
	// - goose
	// - sql-migrate
	//
	// Here we just return nil to indicate success.

	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at INTEGER NOT NULL DEFAULT (strftime('%s','now') * 1000000),
		updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now') * 1000000)
	);
	`

	listingTable := `
	CREATE TABLE IF NOT EXISTS listings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		price INTEGER NOT NULL CHECK(price > 0),
		listing_type TEXT NOT NULL CHECK(listing_type IN ('rent', 'sale')),
		created_at INTEGER NOT NULL DEFAULT (strftime('%s','now') * 1000000),
		updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now') * 1000000),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(userTable)
	if err != nil {
		log.Printf("Error creating 'users' table: %v", err)
		return err
	}

	_, err = db.Exec(listingTable)
	if err != nil {
		log.Printf("Error creating 'listings' table: %v", err)
		return err
	}

	log.Println("Migration successful: users and listings tables created.")
	return nil
}
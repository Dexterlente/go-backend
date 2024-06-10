package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() (*sqlx.DB, error) {
	// Connect to the database
	db, err := connect()
	if err != nil {
		return nil, err
	}

	// Check if the users table exists
	if !tableExists(db, "users") {
		log.Println("Users table does not exist, running migrations...")
		// Run migrations
		if err := runMigrations(db); err != nil {
			return nil, err
		}
		log.Println("Migrations completed successfully!")
	}

	log.Println("Connected to the database successfully!")
	return db, nil
}

func connect() (*sqlx.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	return sqlx.Connect("postgres", databaseURL)
}

func tableExists(db *sqlx.DB, tableName string) bool {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)", tableName)
	if err != nil {
		log.Printf("Error checking if table exists: %v", err)
	}
	return exists
}

func runMigrations(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}
	log.Println("Users table created successfully!")
	return nil
}

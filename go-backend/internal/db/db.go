package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() (*sqlx.DB, error) {
    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        return nil, fmt.Errorf("DATABASE_URL is not set")
    }

    db, err := sqlx.Connect("postgres", databaseURL)
    if err != nil {
        return nil, err
    }
    log.Println("Connected to the database successfully!")
    return db, nil
}

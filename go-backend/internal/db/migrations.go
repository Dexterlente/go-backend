package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// Run all migrations
func runMigrations(db *sqlx.DB) error {
    migrations := []func(*sqlx.DB) error{
        migrateUsers,
        // migrateProducts,  // Example additional table migration
        // Add more migrations here
    }

    for _, migrate := range migrations {
        if err := migrate(db); err != nil {
            return err
        }
    }
    log.Println("Migrations completed successfully!")
    return nil
}

// Migrate Users Table
func migrateUsers(db *sqlx.DB) error {
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

// // Migrate Products Table (Example additional table)
// func migrateProducts(db *sqlx.DB) error {
//     _, err := db.Exec(`
//         CREATE TABLE IF NOT EXISTS products (
//             id SERIAL PRIMARY KEY,
//             name TEXT NOT NULL,
//             price NUMERIC NOT NULL
//         )
//     `)
//     if err != nil {
//         return fmt.Errorf("failed to create products table: %v", err)
//     }
//     log.Println("Products table created successfully!")
//     return nil
// }

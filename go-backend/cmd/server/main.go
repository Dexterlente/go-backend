package main

import (
	"go-backend/internal/db"
	"go-backend/internal/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    database, err := db.New()
    if err != nil {
        log.Fatal(err)
    }

    // Set up routes
    http.HandleFunc("/users", handlers.GetUsers(database))
    http.HandleFunc("/user", handlers.CreateUser(database))

    log.Println("Server is running on port 5050")
    log.Fatal(http.ListenAndServe(":5050", nil))
}

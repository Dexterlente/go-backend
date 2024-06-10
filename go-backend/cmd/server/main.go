package main

import (
	"go-backend/internal/db"
	"go-backend/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

	port := os.Getenv("PORT")
    if port == "" {
        port = "5050" 
    }
	
    // Initialize database
    database, err := db.New()
    if err != nil {
        log.Fatal(err)
    }

    // Set up routes
    http.HandleFunc("/users", handlers.GetUsers(database))
    http.HandleFunc("/user", handlers.CreateUser(database))

	log.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

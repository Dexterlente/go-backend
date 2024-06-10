package repositories

import (
	"go-backend/internal/models"

	"github.com/jmoiron/sqlx"
)

func GetUsers(db *sqlx.DB) ([]models.User, error) {
    var users []models.User
    err := db.Select(&users, "SELECT id, name FROM users")
    return users, err
}

func CreateUser(db *sqlx.DB, user *models.User) (int64, error) {
    result, err := db.Exec("INSERT INTO users (name) VALUES ($1)", user.Name)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

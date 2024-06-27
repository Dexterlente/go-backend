package repositories

import (
	"go-backend/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(db *sqlx.DB) ([]models.User, error) {
    var users []models.User
    err := db.Select(&users, "SELECT id, first_name, last_name, email, username FROM users")
    return users, err
}

func CreateUser(db *sqlx.DB, user *models.User) (int, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return 0, err
    }

    var id int
    err = db.QueryRowx(`INSERT INTO users (first_name, last_name, email, password, username) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
        user.FirstName, user.LastName, user.Email, hashedPassword, user.UserName).Scan(&id)

    if err != nil {
        return 0, err
    }
    return id, nil
}
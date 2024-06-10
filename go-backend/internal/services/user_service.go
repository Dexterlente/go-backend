package services

import (
	"go-backend/internal/models"
	"go-backend/internal/repositories"

	"github.com/jmoiron/sqlx"
)

func GetUsers(db *sqlx.DB) ([]models.User, error) {
    return repositories.GetUsers(db)
}

func CreateUser(db *sqlx.DB, user *models.User) (int64, error) {
    return repositories.CreateUser(db, user)
}

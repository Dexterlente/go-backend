package handlers

import (
	"encoding/json"
	"go-backend/internal/models"
	"go-backend/internal/services"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func GetUsers(db *sqlx.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        users, err := services.GetUsers(db)
        if err != nil {
            ErrorResponse(w, http.StatusInternalServerError, err.Error())
            return
        }
        JSONResponse(w, http.StatusOK, users)
    }
}

func CreateUser(db *sqlx.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            ErrorResponse(w, http.StatusBadRequest, err.Error())
            return
        }

        id, err := services.CreateUser(db, &user)
        if err != nil {
            ErrorResponse(w, http.StatusInternalServerError, err.Error())
            return
        }

        JSONResponse(w, http.StatusCreated, map[string]interface{}{"id": id})
    }
}
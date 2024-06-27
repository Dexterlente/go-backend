package models

import (
	"time"
)
type User struct {
    ID        int       `db:"id" json:"id"`
    FirstName string    `db:"first_name" json:"first_name"`
    LastName  string    `db:"last_name" json:"last_name"`
    Email     string    `db:"email" json:"email"`
    UserName  string    `db:"username" json:"username"`
    Password  string    `db:"-" json:"-"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

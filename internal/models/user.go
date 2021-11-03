package models

import (
	"time"
)

const UsersTable = "users"

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Status bool `json:"status"`
	Config string `json:"config"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

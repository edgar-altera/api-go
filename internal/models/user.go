package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

const UsersTable = "users"

type User struct {
	Id int `json:"id"`
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	PasswordHash string `json:"-"`
	Status bool `json:"status"`
	IsAdmin bool `json:"is_admin"`
	Config string `json:"config"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
    
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    
	if err != nil {
        return err
    }

    u.PasswordHash = string(passwordHash)

    return nil
}

func (u User) PasswordMatch(password string) bool {
    
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

    return err == nil
}

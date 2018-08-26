package models

import "time"

//Roles constants
const (
	GuestRole = "guest"
	UserRole  = "user"
	AdminRole = "admin"
)

type User struct {
	timestampable
	ID                uint64    `json:"id"`
	Username          string    `json:"username" validate:"required"`
	Email             string    `json:"email" validate:"required,email"`
	Name              string    `json:"name"`
	Bio               string    `json:"bio"`
	Role              string    `json:"-" validate:"required"`
	EncryptedPassword string    `json:"-" validate:"required"`
	CreatedAt         time.Time `json:"-"`
}

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
	MediaID           uint64    `json:"media_id"`
	Media             *Media    `json:"-"`
	Role              string    `json:"-" validate:"required"`
	EncryptedPassword string    `json:"-" validate:"required"`
	CreatedAt         time.Time `json:"-"`
}

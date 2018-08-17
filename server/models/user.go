package models

import "time"

type User struct {
	timestampable
	ID                uint64    `json:"id"`
	Username          string    `json:"username" validate:"required"`
	Email             string    `json:"email" validate:"required,email"`
	Name              string    `json:"name"`
	Bio               string    `json:"bio"`
	EncryptedPassword string    `json:"-" validate:"required"`
	CreatedAt         time.Time `json:"-"`
}

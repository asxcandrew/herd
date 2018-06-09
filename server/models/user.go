package models

import "time"

type User struct {
	ID                uint64
	Username          string `validate:"required"`
	Email             string `validate:"required,email"`
	Name              string
	Bio               string
	EncryptedPassword string `validate:"required"`
	CreatedAt         time.Time
}

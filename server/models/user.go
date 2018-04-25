package models

import "time"

type User struct {
	ID                int64
	Username          string `validate:"required"`
	Email             string `validate:"required,email"`
	EncryptedPassword string `validate:"required"`
	CreatedAt         time.Time
}

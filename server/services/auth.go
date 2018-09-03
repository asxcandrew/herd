package services

import (
	"github.com/asxcandrew/herd/server/models"
	"golang.org/x/crypto/bcrypt"
)

//Login service
func Login(email string, password string) (*models.User, error) {
	user := &models.User{}
	err := models.DB().Model(user).Where("email = ?", email).Column("Media").Select()

	if err != nil {
		return nil, err
	}

	byteHash := []byte(user.EncryptedPassword)
	bytePass := []byte(password)

	err = bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//Register service
func Register(user *models.User, password string) error {
	bytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.EncryptedPassword = string(hashedBytes[:])
	user.Role = models.UserRole

	err = models.Create(user)

	if err != nil {
		return err
	}

	return nil
}

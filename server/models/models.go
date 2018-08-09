package models

import (
	"math/rand"
	"time"

	validator "gopkg.in/go-playground/validator.v8"
)

var (
	vldtr   = validator.New(&validator.Config{TagName: "validate"})
	letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

//Validate creates resource with validation
func Validate(model interface{}) error {
	return vldtr.Struct(model)
}

//Create creates models with validation
func Create(model interface{}) error {
	err := Validate(model)

	if err != nil {
		return err
	}

	err = DB().Insert(model)

	if err != nil {
		return err
	}
	return nil
}

//RandString returns rand string by given length
func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

package models

import (
	validator "gopkg.in/go-playground/validator.v8"
)

type FullUserStruct struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

//CreateUser creates user with validation
func (u *User) Create() error {
	validate := validator.New(&validator.Config{TagName: "validate"})
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	err = DB().Insert(u)

	if err != nil {
		return err
	}
	return nil
}

func IsAvailableUsername(username string) (bool, error) {
	count, err := DB().Model((*User)(nil)).Where("username = ?", username).Count()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return true, nil
	}
	return false, nil
}

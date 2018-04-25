package models

import validator "gopkg.in/go-playground/validator.v8"

//CreateUser creates user with validation
func CreateUser(u *User) error {
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

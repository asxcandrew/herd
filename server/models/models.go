package models

import validator "gopkg.in/go-playground/validator.v8"

var vldtr = validator.New(&validator.Config{TagName: "validate"})

//Validate creates resource with validation
func Validate(model interface{}) error {
	return vldtr.Struct(model)
}

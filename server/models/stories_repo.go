package models

//Create creates story with validation
func (u *Story) Create() error {
	err := Validate(u)

	if err != nil {
		return err
	}

	err = DB().Insert(u)

	if err != nil {
		return err
	}
	return nil
}

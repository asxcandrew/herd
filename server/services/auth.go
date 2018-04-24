package services

//Login service
func Login(email string, password string) (string, error) { //should return user model
	// user find by email

	// byteHash := []byte(user.PassworHash)

	// err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	// if err != nil {
	// 	log.Println(err)
	// 	return "", err
	// }
	// return user.Id, nil
	return "", nil
}

//Register service
func Register(email string, password string) (string, error) { //should return user model
	// bytes := []byte(password)
	// hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	// if err != nil {
	// 	return "", err
	// }

	// hash := string(hashedBytes[:])

	// validate user
	// create user here with hash
	return "", nil
}

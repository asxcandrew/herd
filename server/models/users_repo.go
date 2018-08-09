package models

type FullUserStruct struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
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

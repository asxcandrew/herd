package models

type FullUserStruct struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type UserResponse struct {
	Username string `json:"html_body"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Name     string `json:"name"`
	MediaID  uint64 `json:"media_id"`
}

func IsAvailableUsername(username string) bool {
	count, err := DB().Model((*User)(nil)).Where("username = ?", username).Count()
	if err != nil {
		return false
	}

	return (count == 0)
}

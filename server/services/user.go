package services

import (
	"fmt"

	"github.com/asxcandrew/herd/server/models"
)

func UpdateUser(user *models.User, params *models.UserResponse) error {
	if params.MediaID != 0 {
		DeleteAvatar(user)
		user.MediaID = params.MediaID
	}

	fmt.Println(params, user)

	err := models.Update(user)
	return err
}

//DeleteAvatar to delete media object from db and storage
func DeleteAvatar(user *models.User) error {
	if user.Media != nil {
		err := DeleteMedia(user.Media)

		if err != nil {
			return err
		}

		err = models.DB().Delete(user.Media)
		return err
	}
	return nil
}

package services

import "github.com/asxcandrew/herd/server/models"

func DeleteMedia(media *models.Media) error {
	err := Storage().RemoveObject(media.Name)
	return err
}

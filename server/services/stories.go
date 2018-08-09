package services

import (
	"strings"

	"github.com/asxcandrew/herd/server/models"
)

const (
	maxTitleLength = 50
)

var replacer = strings.NewReplacer(" ", "-", "\t", "-", "`", "-")

//CreateStory creates story with default values
func CreateStory(story *models.Story) error {
	story.UID = models.RandString(10)
	if story.Link == "" {
		pickLength := maxTitleLength
		if len(story.Title) < maxTitleLength {
			pickLength = len(story.Title)
		}
		story.Link = replacer.Replace(strings.ToLower(story.Title))[0:pickLength]
	}

	err := models.Create(story)

	return err
}

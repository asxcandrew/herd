package services

import (
	"strconv"
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

func UpdateStory(story *models.Story, params *models.StoryResponse) error {
	if params.Title != "" {
		story.Title = params.Title
	}
	if params.HTMLBody != "" {
		story.HTMLBody = params.HTMLBody
	}

	if params.Active != "" {
		isActive, err := strconv.ParseBool(params.Active)

		if err != nil {
			return err
		}
		story.Active = isActive
	}

	if story.PublishedAt.IsZero() && !params.PublishedAt.IsZero() {
		story.PublishedAt = params.PublishedAt
	}

	err := models.Update(story)

	return err
}

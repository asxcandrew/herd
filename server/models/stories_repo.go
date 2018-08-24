package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type StoryResponse struct {
	HTMLBody    string    `json:"html_body"`
	Title       string    `json:"title"`
	Active      string    `json:"active"`
	PublishedAt time.Time `json:"published_at"`
}

//QueryStories returns query for index path
func QueryStories(stories *[]Story) *orm.Query {
	return DB().Model(stories).Column("Author", "Tags")
}

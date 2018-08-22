package models

import "time"

type StoryResponse struct {
	HTMLBody    string    `json:"html_body"`
	Title       string    `json:"title"`
	Active      string    `json:"active"`
	PublishedAt time.Time `json:"published_at"`
}

//QueryStories returns query for index path
func QueryStories(stories *[]Story) error {
	return DB().Model(stories).Column("Author", "Tags").Select()
}

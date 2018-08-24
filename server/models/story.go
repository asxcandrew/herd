package models

import "time"

// Story db structure
type Story struct {
	timestampable

	ID          uint64    `json:"-"`
	UID         string    `json:"uid" validate:"required"`
	Link        string    `json:"link"`
	HTMLBody    string    `json:"-" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Subtitle    string    `json:"subtitle"`
	AuthorID    uint64    `json:"-" validate:"required"`
	Author      *User     `json:"user"`
	Tags        []Tag     `json:"tags" pg:"many2many:story_tags"`
	Active      bool      `json:"active"`
	PublishedAt time.Time `json:"published_at"`
}

// StoryTags structure to keep related tags
type StoryTags struct {
	StoryID uint64
	TagID   uint64
}

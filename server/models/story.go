package models

import "time"

// Story db structure
type Story struct {
	ID          uint64
	UID         string
	Link        string
	Title       string
	Subtitle    string
	AuthorID    uint64 `validate:"required"`
	Revisions   []*StoryRevision
	Author      *User
	Tags        []Tag `pg:"many2many:story_tags"`
	Active      bool
	PublishedAt time.Time
	CreatedAt   time.Time
}

// StoryRevision structure to keep story revisions
type StoryRevision struct {
	ID        uint64
	Markdown  string
	StoryID   uint64 `validate:"required"`
	Story     *Story
	CreatedAt time.Time
}

// StoryTags structure to keep related tags
type StoryTags struct {
	StoryID uint64
	TagID   uint64
}

package models

// Topic db structure
type Topic struct {
	ID      uint64
	Name    string `validate:"required"`
	URI     string `validate:"required"`
	Caption string `validate:"required"`
	Active  bool
	Related []Topic `pg:"many2many:related_topics,joinFK:related_id"`
}

// RelatedTopics struct to connect topics together
type RelatedTopics struct {
	RelatedID uint64
	TopicID   uint64
}

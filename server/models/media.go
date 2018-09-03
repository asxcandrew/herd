package models

// Media db structure
type Media struct {
	timestampable
	ID          uint64 `json:"id"`
	ContentType string `json:"content_type" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

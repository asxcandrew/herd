package models

// Tag db structure
type Tag struct {
	ID    uint64
	Name  string `validate:"required"`
	Count uint64
}

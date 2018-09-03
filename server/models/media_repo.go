package models

type MediaResponse struct {
	ID          uint64 `json:"id"`
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
}

package models

type URLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortenedURL string `json:"bitly_url"`
}

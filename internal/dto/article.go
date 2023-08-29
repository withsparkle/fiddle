package dto

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID      uuid.UUID `json:"id"`
	URL     *url.URL  `json:"url"`
	Tags    []string  `json:"tags"`
	Links   []Link    `json:"taxonomy"`
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Content string    `json:"body"`
	Summary string    `json:"summary"`
}

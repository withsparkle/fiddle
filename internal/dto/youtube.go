package dto

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type YouTubeVideo struct {
	ID      uuid.UUID `json:"id"`
	URL     *url.URL  `json:"url"`
	Tags    []string  `json:"tags"`
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Desc    string    `json:"body"`
	Summary string    `json:"summary"`
}

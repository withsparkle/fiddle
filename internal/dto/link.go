package dto

import (
	"net/url"

	"github.com/google/uuid"
)

type LinkType string

const (
	Category  LinkType = "category"
	Taxonomy  LinkType = "taxonomy"
	Reference LinkType = "reference"
)

type Link struct {
	ID   uuid.UUID `json:"id"`
	URL  *url.URL  `json:"url"`
	Type LinkType  `json:"kind"`
	Name string    `json:"name"`
}

package models

import (
	"time"
)

type DocumentResponse struct {
	ID        string    `string:"id"`
	Title     string    `json:"title"`
	Type string `json:"type"`
	Authors []string `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsHosted bool `json:"is_hosted"`
	ExternalLink string `json:"external_link"`
	Tags []string `json:"tags"`
	ReferencedDocuments []string `json:"referenced"`
}

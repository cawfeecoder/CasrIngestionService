package models

type DocumentRequest struct {
	Title string `json:"title"`
	Type string `json:"type"`
	Content string `json:"content"`
	IsHosted bool `json:"is_hosted"`
	ExternalLink string `json:"external_link"`
	Tags []string `json:"tags"`
}
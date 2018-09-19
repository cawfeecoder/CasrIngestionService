package models

import "time"

type Document struct {
	ID string `json:"id"` //UUID corresponding to the ES document
	Title string `json:"title"`
	Authors []string `json:"author"` //Author(s) of the document
	Type string `json:"type"` //Type can take in two values: file, document. It let's us know whether we're gonna resolve the external link
	Content string `json:"content"` //In both cases, we display the content of the document
	CreatedAt time.Time `json:"created_at"`
	IsHosted bool `json:"is_hosted"` //Is the file hosted on our servers?
	ExternalLink string `json:"external_link"` //Points to either a storage block on our server or externally
	Tags []string `json:"tags"` //Tag the document
	ReferencedDocuments []string `json:"referenced"` //Any referenced documents
}

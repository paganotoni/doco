package internal

import (
	"fmt"
	"html/template"
	"path/filepath"
)

// document represents a single markdown document in the site
// it contains the metadata and the content of the document.
type document struct {
	filename string
	title    string

	// content
	markdown []byte
	html     template.HTML
}

func (doc document) String() string {
	return fmt.Sprintf("Document: %v", doc.title)
}

// NewDocument takes the path of a document and its content
// and returns a parsed document with the metadata applied.
func NewDocument(path string, content []byte) (document, error) {
	// Parse the metadata and apply it to the document
	meta, err := metadataFrom(content)
	if err != nil {
		return document{}, fmt.Errorf("error parsing metadata for %v: %w", path, err)
	}

	html, err := htmlFrom(content)
	if err != nil {
		return document{}, fmt.Errorf("error generating html for %v: %w", path, err)
	}

	title, ok := meta["title"].(string)
	if !ok {
		// Use filename as title
		title = humanizeFilename(path)
	}

	doc := document{
		title:    title,
		filename: filepath.Base(path),

		html:     html,
		markdown: content,
	}

	return doc, nil
}

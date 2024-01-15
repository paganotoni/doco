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
	index    int

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

	index, ok := meta["index"].(int)
	if !ok {
		index = 10_000_000
	}

	doc := document{
		title:    title,
		filename: filepath.Base(path),
		index:    index,

		html:     html,
		markdown: content,
	}

	return doc, nil
}

type documents []document

func (d documents) Len() int           { return len(d) }
func (d documents) Less(i, j int) bool { return d[i].index < d[j].index }
func (d documents) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

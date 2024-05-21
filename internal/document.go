package internal

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/paganotoni/doco/internal/markdown"
	"golang.org/x/net/html"
)

// document represents a single markdown document in the site
// it contains the metadata and the content of the document.
type document struct {
	filename    string
	title       string
	index       int
	description string
	keywords    string

	// content
	markdown []byte
	html     template.HTML
}

func (doc document) String() string {
	return fmt.Sprintf("Document: %v", doc.title)
}

func (doc document) Tokens() string {
	var s string

	domDocTest := html.NewTokenizer(strings.NewReader(string(doc.html)))
	previousStartTokenTest := domDocTest.Token()
l:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break l
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}

			content := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			if len(content) > 0 {
				s += content + " "
			}
		}
	}

	return s
}

// NewDocument takes the path of a document and its content
// and returns a parsed document with the metadata applied.
func NewDocument(path string, content []byte) (document, error) {
	// Parse the metadata and apply it to the document
	meta, err := markdown.ReadMetadata(content)
	if err != nil {
		return document{}, fmt.Errorf("error parsing metadata for %v: %w", path, err)
	}

	html, err := markdown.HTMLFrom(content)
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

	description, ok := meta["description"].(string)
	if !ok {
		// Use filename as title
		description = ""
	}

	keywords, ok := meta["keywords"].(string)
	if !ok {
		keywords = ""
	}

	doc := document{
		title:    title,
		filename: filepath.Base(path),
		index:    index,

		description: description,
		keywords:    keywords,

		html:     html,
		markdown: content,
	}

	return doc, nil
}

type documents []document

func (d documents) Len() int           { return len(d) }
func (d documents) Less(i, j int) bool { return d[i].index < d[j].index }
func (d documents) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

package internal

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/paganotoni/doco/internal/markdown"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NewDocument takes the path of a document and its content
// and returns a parsed document with the metadata applied.
func NewDocument(path string, content []byte) (document, error) {
	// Parse the metadata and apply it to the document
	meta, err := markdown.ReadMetadata(content)
	if err != nil {
		return document{}, fmt.Errorf("error parsing metadata for %v: %w", path, err)
	}

	title, ok := meta["title"].(string)
	if !ok {
		// Humanizing the filename into a title
		title := strings.ReplaceAll(path, filepath.Ext(path), "") // Use filename as title
		title = filepath.Base(title)                              // remove the path
		title = strings.ReplaceAll(title, "-", " ")               // remove dashes
		title = strings.ReplaceAll(title, "_", " ")               // remove underscores
		title = cases.Title(language.English).String(title)
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
		filename: filepath.Base(path),

		index:       index,
		Title:       title,
		description: description,
		keywords:    keywords,
		markdown:    content,
	}

	return doc, nil
}

// document represents a single markdown document in the site
// it contains the metadata and the content of the document.
type document struct {
	filename    string
	Title       string
	index       int
	description string
	keywords    string

	// content
	markdown []byte
	section  *section

	prev *document
	next *document
}

func (doc document) String() string {
	return fmt.Sprintf(
		"Document: %s",
		strings.Join([]string{strings.TrimSpace(doc.section.Name), doc.Title}, " - "),
	)
}

func (doc document) Tokens() (tokens []string) {
	r := map[string]bool{}

	for _, v := range strings.Fields(string(doc.markdown)) {
		if r[v] == true {
			continue
		}

		r[v] = true
		tokens = append(tokens, v)
	}

	return tokens
}

func (doc document) Link() string {
	return path.Join(
		doc.section.path,
		strings.TrimSuffix(doc.filename, ".md")+".html",
	)
}

func (doc document) FileName() string {
	return path.Join(
		doc.section.path,
		strings.TrimSuffix(doc.filename, ".md")+".html",
	)
}

type documents []document

func (d documents) Len() int           { return len(d) }
func (d documents) Less(i, j int) bool { return d[i].index < d[j].index }
func (d documents) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

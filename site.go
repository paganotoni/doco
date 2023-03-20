package doco

import (
	"bytes"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Site being parsed and built from the
// docs markdown files.
type Site struct {
	// Meta data from the _index.md file
	meta map[string]interface{}

	// List of documents parsed from the docs folder
	// files starting with underscore are not parsed
	documents Documents
}

// Takes reads the name of the site from the _index.md
func (s *Site) Name() string {
	name, ok := s.meta["Name"].(string)
	if !ok {
		return ""
	}

	return name
}

// add a document to the site. Useful when walking
// a directory to add the files to the site.
func (s *Site) add(d Document) {
	s.documents = append(s.documents, d)
}

func (s *Site) parseMeta() {
	index := s.documents.Index()
	if index == nil {
		return
	}

	content, err := index.ReadContent()
	if err != nil {
		return
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return
	}

	s.meta = meta.Get(context)
}

func (s *Site) Documents() Documents {
	return s.documents
}

func NewSite(docs Documents) *Site {
	site := &Site{
		meta:      make(map[string]interface{}),
		documents: docs,
	}

	site.parseMeta()

	return site
}

package internal

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var (
	// mparser is the markdown parser used to parse the content
	mparser = goldmark.New(goldmark.WithExtensions(meta.Meta))
)

// metadataFrom parses the content and returns the metadata
// from the markdown file.
func metadataFrom(content []byte) (map[string]interface{}, error) {
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := mparser.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	return meta.Get(context), nil
}

// Generates HTML from the passed content.
func htmlFrom(content []byte) (template.HTML, error) {
	var buf bytes.Buffer
	if err := mparser.Convert(content, &buf); err != nil {
		return template.HTML(""), err
	}

	return template.HTML(buf.String()), nil
}

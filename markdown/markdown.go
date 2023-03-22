// markdown package contains common markdown parsing functions
// abstracting this package allows us to change the markdown
// engine afterwards as long as it covers the same functionality
// for the rest of the packages.
package markdown

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var (
	mparser = goldmark.New(goldmark.WithExtensions(meta.Meta))
)

// MetadataFrom parses the content and returns the metadata
// from the markdown file.
func MetadataFrom(content []byte) (map[string]interface{}, error) {
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := mparser.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	return meta.Get(context), nil
}

// Generates HTML from the passed content.
func HTMLFrom(content []byte) (template.HTML, error) {
	var buf bytes.Buffer
	if err := mparser.Convert(content, &buf); err != nil {
		return template.HTML(""), err
	}

	return template.HTML(buf.String()), nil
}

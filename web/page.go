package web

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/paganotoni/doco"
	"github.com/paganotoni/doco/markdown"
)

type Page struct {
	// SourcePath is the path to the original markdown file.
	// its used to generate the resulting HTML file.
	SourcePath string

	// Metadata from the markdown file.
	Metadata map[string]interface{}

	// HTML generated from the markdown file.
	HTML template.HTML
}

func (p Page) title() string {
	t, ok := p.Metadata["Title"].(string)
	if ok {
		return t
	}

	return p.SourcePath
}

func (p Page) weight() int {
	t, ok := p.Metadata["Weight"].(int)
	if ok {
		return t
	}

	return 0
}

func (p Page) resultPath() (path string) {
	path = strings.Replace(p.SourcePath, "docs/", "", 1)
	path = strings.Replace(path, ".md", ".html", 1)
	path = filepath.Join("public", path)

	return path
}

func (p Page) resultLink() (path string) {
	path = strings.Replace(p.SourcePath, "docs/", "", 1)
	path = strings.Replace(path, ".md", ".html", 1)
	path = filepath.Join(path)

	return path
}

type Pages []Page

// NewPage creates a new Page from the passed document, it
// parses the document content to extract the metadata and generate
// the HTML content.
func NewPage(document doco.Document) (page Page, err error) {
	content, err := document.ReadContent()
	if err != nil {
		return page, fmt.Errorf("error reading document content:%w", err)
	}

	meta, err := markdown.MetadataFrom(content)
	if err != nil {
		return page, fmt.Errorf("error reading document metadata:%w", err)
	}

	htmlc, err := markdown.HTMLFrom(content)
	if err != nil {
		return page, fmt.Errorf("error generating document HTML:%w", err)
	}

	return Page{
		SourcePath: document.Path(),
		Metadata:   meta,
		HTML:       htmlc,
	}, nil
}

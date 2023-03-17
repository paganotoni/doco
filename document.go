package doco

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
)

// Document is a type that represents a documentation page
// it will be used to generate the HTML file and the resulting path
type Document string

func (d Document) ResultingPath() string {
	file := string(d)
	file = strings.TrimPrefix(file, "docs/")
	file = strings.Replace(file, filepath.Ext(file), ".html", 1)

	return filepath.Join("public", file)
}

func (d Document) Href() string {
	file := string(d)
	file = strings.TrimPrefix(file, "docs/")
	file = strings.Replace(file, filepath.Ext(file), ".html", 1)

	return filepath.Join("/", file)
}

func (d Document) LinkName() string {
	file := filepath.Base(string(d))
	return strings.Replace(file, filepath.Ext(file), "", 1)
}

func (d Document) HTML() ([]byte, error) {
	source, err := os.ReadFile(string(d))
	if err != nil {
		return nil, err
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	buf := bytes.Buffer{}
	if err = markdown.Convert(source, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type Documents []Document

func (d Documents) NavigationHTML() string {
	var html string
	for _, doc := range d {
		if string(doc) == "docs/index.md" {
			continue
		}

		html += fmt.Sprintf(`<li><a class="block px-4 py-2" href="%s">%s</a></li>`, doc.Href(), doc.LinkName())
	}

	return html
}

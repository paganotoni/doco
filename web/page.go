package web

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/paganotoni/doco"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var markdown = goldmark.New(goldmark.WithExtensions(meta.Meta))

type Page struct {
	// Meta data from the markdown file.
	meta map[string]interface{}

	// Site being built
	Site *doco.Site

	// Document being rendered
	Document doco.Document

	// The content of the document
	Content template.HTML
}

func (p Page) SiteTitle() string {
	return p.Site.Name()
}

func (p Page) Title() string {
	title, ok := p.meta["Title"].(string)
	if !ok {
		return ""
	}

	return title
}

func (p Page) HTML() (template.HTML, error) {
	content, err := p.Document.ReadContent()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := markdown.Convert(content, &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

func (p Page) Navigation() template.HTML {
	items := ``
	for _, d := range p.Site.Documents() {
		path := strings.Replace(d.Path(), "docs/", "", 1)
		path = strings.Replace(path, ".md", ".html", 1)

		p, err := NewPage(p.Site, d)
		if err != nil {
			continue
		}

		items += fmt.Sprintf(`<li class="nav-item"><a href="%s">%s</a></li>`, path, p.Title())
	}

	return template.HTML(items)
}

func NewPage(site *doco.Site, document doco.Document) (Page, error) {
	p := Page{
		Site:     site,
		Document: document,
	}

	content, err := document.ReadContent()
	if err != nil {
		return p, err
	}

	context := parser.NewContext()
	err = markdown.Convert(content, io.Discard, parser.WithContext(context))
	if err != nil {
		return p, err
	}

	p.meta = meta.Get(context)

	return p, nil
}

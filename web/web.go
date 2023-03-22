// web package is used to build the web interface
// for the documentation after the parsing has been
// done by the doco package. It uses doco.Site and
// doco.Document to build the web interface.
package web

import (
	_ "embed"
	"fmt"
	"os"

	"html/template"

	"github.com/paganotoni/doco"
)

var (
	//go:embed templates/page.html
	pgTmpl string

	//The template/html template used to generate each of the web pages.
	pageTemplate = template.Must(template.New("page").Parse(pgTmpl))
)

type pageTemplateData struct {
	SiteTitle  string
	HTML       template.HTML
	Navigation template.HTML
}

// Generate the web interface for the documentation.
func Generate(site *doco.Site) error {
	// TODO: Copy Scripts (Mobile menu, Search)
	// Write the style.css file
	writeStyles()

	pages := Pages{}
	for _, page := range site.Documents() {
		p, err := NewPage(page)
		if err != nil {
			return err
		}

		pages = append(pages, p)
	}

	// Generate the document pages
	for _, page := range pages {
		// transforming the path to the public folder
		path := page.resultPath()
		fx, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err
		}

		err = pageTemplate.Execute(fx, pageTemplateData{
			SiteTitle:  "TODO",
			HTML:       page.HTML,
			Navigation: navigation(pages).HTMLFor(page),
		})

		if err != nil {
			return err
		}

		fmt.Printf("[INFO] Write > %s\n", path)
	}

	return nil
}

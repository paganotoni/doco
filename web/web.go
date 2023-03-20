// web package is used to build the web interface
// for the documentation after the parsing has been
// done by the doco package. It uses doco.Site and
// doco.Document to build the web interface.
package web

import (
	_ "embed"

	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/paganotoni/doco"
)

var (
	//go:embed templates/page.html
	pageTemplate string
)

// Generate the web interface for the documentation.
func Generate(site *doco.Site) error {
	// Write the style.css file
	writeStyles()

	// Parse the page template
	tt := template.New("page")
	tmpl := template.Must(tt.Parse(pageTemplate))

	// Generate the index page
	fmt.Println("[INFO] Write > public/index.html")

	// Generate the document pages
	for _, page := range site.Documents() {
		// transforming the path to the public folder
		path := strings.Replace(page.Path(), "docs/", "", 1)
		path = strings.Replace(path, ".md", ".html", 1)
		path = filepath.Join("public", path)

		fx, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err
		}

		fmt.Printf("[INFO] Write > %s\n", path)
		page, err := NewPage(site, page)
		if err != nil {
			return err
		}

		err = tmpl.Execute(fx, page)
		if err != nil {
			return err
		}
	}

	return nil

	// TODO: Copy Scripts (Mobile menu, Search)
}

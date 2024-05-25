package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/paganotoni/doco/internal/config"
	"github.com/paganotoni/doco/internal/markdown"
)

var (
	//go:embed assets/doco.css
	style template.CSS

	//go:embed assets/doco.js
	docoJS template.JS

	//go:embed page.html
	pageHTML string
	pageTmpl = template.Must(
		template.New("page").Funcs(template.FuncMap{
			"htmlFrom": func(m []byte) template.HTML {
				c, err := markdown.HTMLFrom(m)
				if err != nil {
					return ""
				}
				return template.HTML(c)
			},
		}).Parse(pageHTML),
	)
)

// Generates the static html files for the site
// and writes them to the destination folder.
func Generate(srcFolder, destination string, site *site) error {
	// Cleanup the folder
	err := os.RemoveAll(destination)
	if err != nil {
		return err
	}

	conf, err := config.Read(srcFolder)
	if err != nil {
		return err
	}

	// Generate pages for each of the sections and documents inside them
	// and write them to the destination folder.
	for _, v := range site.sections {
		err := os.MkdirAll(filepath.Join(destination, v.path), os.ModePerm)
		if err != nil {
			return err
		}

		for _, doc := range v.documents {
			bb := bytes.NewBuffer([]byte{})
			err = pageTmpl.Execute(bb, struct {
				Site config.Site

				Title       string
				Name        string
				SectionName string
				Description string
				Keywords    string

				NextLink  string
				NextTitle string
				PrevLink  string
				PrevTitle string

				Markdown []byte
				Style    template.CSS
				JS       template.JS
			}{
				Site: conf,

				Name:        "", //doc.Name(),
				Title:       doc.title,
				SectionName: doc.section.name,
				Markdown:    doc.markdown,
				Style:       style,
				JS:          docoJS,
			})

			if err != nil {
				return err
			}

			f := filepath.Join(destination, doc.FileName())
			err = os.WriteFile(f, bb.Bytes(), 0644)
			if err != nil {
				return err
			}
		}
	}

	// Copy assets folder to the destination folder.
	err = copyDir(filepath.Join(srcFolder, "assets"), filepath.Join(destination, "assets"))
	if err != nil {
		return fmt.Errorf("error copying assets: %w", err)
	}

	// // Generating the site index file to be used by the search.
	// f, err := os.Create(filepath.Join(destination, "index.json"))
	// if err != nil {
	// 	return fmt.Errorf("error generating search index: %w", err)
	// }

	// encoder := json.NewEncoder(f)
	// err = encoder.Encode(pages)
	// if err != nil {
	// 	return fmt.Errorf("error generating search index: %w", err)
	// }

	return nil
}

// generatedPage is the data passed to the template
// to generate the static html files.
// type generatedPage struct {
// 	filePath string `json:"-"`
// 	Prev     struct {
// 		Title string
// 		Link  string
// 	} `json:"-"`

// 	Next struct {
// 		Title string
// 		Link  string
// 	} `json:"-"`

// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Keywords    string `json:"keywords"`
// 	SectionName string `json:"section_name"`
// 	Link        string `json:"link"`
// 	Tokens      string `json:"content"`

// 	Content template.HTML `json:"-"`
// }

// NAV().CLASS("documents").Children(
// 		Range(s.sections, func(s section) ElementRenderer {
// 			return SECTION().IfChildren(s.name != "", H3().Text(s.name)).Children(
// 				UL().Children(
// 					Range(s.documents, func(d document) ElementRenderer {
// 						link := "/" + filepath.Join(s.path, strings.TrimSuffix(d.filename, ".md")+".html")
// 						return LI().IfCLASS(doc.filename == d.filename, "active").Children(
// 							A().HREF(link).Text(d.title),
// 						)
// 					}),
// 				),
// 			)
// 		}),
// 	)

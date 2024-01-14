package internal

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed style.css
	style []byte

	//go:embed template.html
	tmplHTML []byte
	tmpl     = template.Must(template.New("doco").Parse(string(tmplHTML)))
)

// Generates the static html files for the site
// and writes them to the destination folder.
func Generate(srcFolder, dstFolder string, site *site) error {
	// Cleanup the folder
	err := os.RemoveAll(dstFolder)
	if err != nil {
		return err
	}

	// Create the folder
	err = os.MkdirAll(dstFolder, os.ModePerm)
	if err != nil {
		return err
	}

	siteConfig, err := parseMeta(srcFolder)
	if err != nil {
		return err
	}

	// Copy all assets
	err = copyDir(filepath.Join(srcFolder, "assets"), filepath.Join(dstFolder, "assets"))
	if err != nil {
		return fmt.Errorf("error copying assets: %w", err)
	}

	for _, v := range site.sections {
		err := os.MkdirAll(filepath.Join(dstFolder, v.path), os.ModePerm)
		if err != nil {
			return err
		}

		for _, doc := range v.documents {
			// normalize the filename
			name := strings.Replace(doc.filename, filepath.Ext(doc.filename), ".html", 1)
			name = underscore(name)

			// write the file
			file, err := os.Create(filepath.Join(dstFolder, v.path, name))
			if err != nil {
				return err
			}

			data := struct {
				SiteConfig config

				Title       string
				SectionName string
				Content     template.HTML
				Style       template.CSS

				DesktopNavigation template.HTML
			}{
				SiteConfig: siteConfig,

				Title:       doc.title,
				SectionName: v.name,

				Content:           doc.html,
				Style:             template.CSS(style),
				DesktopNavigation: desktopNavigation(site, doc),
			}

			err = tmpl.Execute(file, data)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

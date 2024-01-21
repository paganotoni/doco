package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed style.css
	style []byte

	//go:embed search.js
	searchJSON []byte

	//go:embed template.html
	tmplHTML []byte
	tmpl     = template.Must(template.New("doco").Parse(string(tmplHTML)))
)

type navlink struct {
	Title string `json:"-"`
	Link  string `json:"-"`
}

// generatedPage is the data passed to the template
// to generate the static html files.
type generatedPage struct {
	SiteConfig config `json:"-"`

	Title       string        `json:"title"`
	SectionName string        `json:"section_name"`
	Content     template.HTML `json:"-"`
	Link        string        `json:"link"`
	filePath    string        `json:"-"`
	Tokens      string        `json:"content"`

	Prev       navlink       `json:"-"`
	Next       navlink       `json:"-"`
	Style      template.CSS  `json:"-"`
	Navigation template.HTML `json:"-"`
}

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

	var pages []generatedPage
	// Generate pages for each of the sections and documents inside them
	// and write them to the destination folder.
	for _, v := range site.sections {
		err := os.MkdirAll(filepath.Join(dstFolder, v.path), os.ModePerm)
		if err != nil {
			return err
		}

		for _, doc := range v.documents {
			// normalize the filename
			name := strings.Replace(doc.filename, filepath.Ext(doc.filename), ".html", 1)
			name = underscore(name)

			data := generatedPage{
				filePath:   filepath.Join(dstFolder, v.path, name),
				SiteConfig: siteConfig,

				Title:       doc.title,
				SectionName: v.name,
				Link:        filepath.Join(v.path, name),

				Content: doc.html,
				Style:   template.CSS(style),
				Tokens:  doc.Tokens(),

				Navigation: desktopNavigation(site, doc),
			}

			pages = append(pages, data)
		}
	}

	// Generate all of the files

	for index, v := range pages {
		if index < len(pages)-1 {
			v.Next.Link = pages[index+1].Link
			v.Next.Title = pages[index+1].Title
		}

		if index > 0 {
			v.Prev.Link = pages[index-1].Link
			v.Prev.Title = pages[index-1].Title
		}

		// write the file
		file, err := os.Create(v.filePath)
		if err != nil {
			return err
		}

		err = tmpl.Execute(file, v)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filepath.Join(dstFolder, "index.json"))
	if err != nil {
		return fmt.Errorf("error generating search index: %w", err)
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(pages)
	if err != nil {
		return fmt.Errorf("error generating search index: %w", err)
	}

	err = os.WriteFile(filepath.Join(dstFolder, "search.js"), searchJSON, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing search js: %w", err)
	}

	return nil
}

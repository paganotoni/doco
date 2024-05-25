package internal

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
			// htmlFrom generates html from markdown
			// this is useful for the content of the page
			// to be generated.
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
func Generate(srcFolder, destination string, s *site) error {
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
	for _, v := range s.Sections {
		err := os.MkdirAll(filepath.Join(destination, v.path), os.ModePerm)
		if err != nil {
			return err
		}

		type docData struct {
			Config config.Site
			Site   *site

			Title       string
			SectionName string
			Description string
			Keywords    string
			Link        string

			Markdown []byte
			Style    template.CSS
			JS       template.JS

			NextLink  string
			NextTitle string
			PrevLink  string
			PrevTitle string
		}

		for _, doc := range v.Documents {
			d := docData{
				Config: conf,
				Site:   s,
				Style:  style,
				JS:     docoJS,

				Title:       doc.Title,
				SectionName: doc.section.Name,
				Markdown:    doc.markdown,
				Link:        doc.Link(),
			}

			if doc.next != nil {
				d.NextLink = doc.next.Link()
				d.NextTitle = doc.next.Title
			}

			if doc.prev != nil {
				d.PrevLink = doc.prev.Link()
				d.PrevTitle = doc.prev.Title
			}

			bb := bytes.NewBuffer([]byte{})
			err = pageTmpl.Execute(bb, d)
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

	// Generating the site index file to be used by the search.
	f, err := os.Create(filepath.Join(destination, "index.json"))
	if err != nil {
		return fmt.Errorf("error generating search index: %w", err)
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(s.SearchData())
	if err != nil {
		return fmt.Errorf("error generating search index: %w", err)
	}

	return nil
}

// copyDir copies a directory recursively from src to dst
// src and dst must be absolute paths. This is useful to copy the
// assets recursively.
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, path[len(src):])
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		return os.Chmod(dstPath, info.Mode())
	})
}

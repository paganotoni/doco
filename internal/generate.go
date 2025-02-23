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
	"strings"

	"github.com/yuin/goldmark"

	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
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
			"htmlFrom": htmlFromMarkdown,
		}).Parse(pageHTML),
	)

	// mparser is the markdown parser used to parse the content
	mparser = goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // read note
		),

		goldmark.WithExtensions(
			// anchor is used to generate anchor links for headings
			// in the markdown file.
			&anchor.Extender{
				Texter: anchor.Text("#"),
				Attributer: anchor.Attributes{
					"class": "heading-anchor",
					"alt":   "Link to this section",
				},
			},
		),
	)
)

// Generates the static html files for the site
// and writes them to the destination folder.
func Generate(srcFolder, destination string, s *site, conf siteConfig) error {
	// Cleanup the folder
	err := os.RemoveAll(destination)
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
			Config siteConfig
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
			err = os.WriteFile(f, bb.Bytes(), 0o644)
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

// htmlFromMarkdown generates html from markdown
// this is useful for the content of the page
// to be generated.
func htmlFromMarkdown(m []byte) template.HTML {
	// Remove the front matter meta from the markdown
	content := string(m)
	content = strings.Replace(content, "---", "", 1)
	c := strings.Index(string(content), "---")
	if c != -1 {
		content = content[c+3:]
	}

	// Convert the markdown to html
	var buf bytes.Buffer
	if err := mparser.Convert([]byte(content), &buf); err != nil {
		return template.HTML("")
	}

	return template.HTML(buf.String())
}

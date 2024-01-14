package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// site represents the documentation site
// it contains all the sections and documents
// that are used to generate the static html files.
type site struct {
	Name string

	sections sections
}

func (s *site) String() string {
	pp := "Site: \n"
	for _, sec := range s.sections {
		pp += "   Section: " + sec.String()
	}

	return pp
}

func (s *site) Add(path string, doc document) error {
	secName := humanize(filepath.Base(filepath.Dir(path)))
	for i, v := range s.sections {
		if v.name == secName {
			v.documents = append(v.documents, doc)
			s.sections[i] = v
			return nil
		}
	}

	sec := section{
		name:      secName,
		path:      filepath.Dir(path),
		documents: []document{doc},
	}

	s.sections = append(s.sections, sec)
	return nil
}

// Reads the folder and returns the parsed site with all the documents
// this site will be used to generate the static html files.
func NewSite(folder string) (*site, error) {
	site := &site{
		sections: sections{},
	}

	// Parse metadata

	err := filepath.Walk(folder, func(path string, d os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories, files that don't have the .md extension
		// and files that start with an underscore.
		if d.IsDir() || filepath.Ext(path) != ".md" || strings.HasPrefix(filepath.Base(path), "_") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		bb, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		path = strings.TrimPrefix(path, folder+"/")
		doc, err := NewDocument(path, bb)
		if err != nil {
			return fmt.Errorf("error parsing %v: %w", path, err)
		}

		return site.Add(path, doc)
	})

	return site, err
}

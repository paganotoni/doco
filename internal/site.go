package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/paganotoni/doco/internal/markdown"
)

// Reads the folder and returns the parsed site with all the documents
// this site will be used to generate the static html files.
func NewSite(folder string) (*site, error) {
	site := &site{
		sections: sections{},
	}

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

	// Adding indexes to the sections
	for i, v := range site.sections {
		metaFile := filepath.Join(folder, v.path, "_meta.md")

		f, err := os.Open(metaFile)
		if err == nil {
			bb, err := io.ReadAll(f)
			if err != nil {
				continue
			}

			meta, err := markdown.ReadMetadata(bb)
			if err == nil {
				var ok bool
				site.sections[i].index, ok = meta["index"].(int)
				if !ok {
					// 10 million to make sure it is the last one by default
					v.index = 10_000_000
				}
			}
		}
	}

	// Sort site secitons by index
	sort.Sort(site.sections)

	// Sort documents
	for i := range site.sections {
		sort.Sort(site.sections[i].documents)
	}

	return site, err
}

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
	// Cover the root case by setting the section name to
	// an empty string
	if secName == "." {
		secName = ""
	}

	for i, v := range s.sections {
		if v.name == secName {
			doc.section = &v
			v.documents = append(v.documents, doc)
			s.sections[i] = v

			return nil
		}
	}

	sec := section{
		name: secName,
		path: filepath.Dir(path),
	}

	doc.section = &sec
	sec.documents = append(sec.documents, doc)

	s.sections = append(s.sections, sec)
	return nil
}

// A section is a group of documents that are in the same folder.
// The section name is the folder name.
type section struct {
	name  string
	path  string
	index int

	documents documents
}

func (s *section) String() string {
	pp := s.name + "\n"
	for _, doc := range s.documents {
		pp += "      " + doc.String() + "\n"
	}

	return pp
}

type sections []section

func (a sections) Len() int           { return len(a) }
func (a sections) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sections) Less(i, j int) bool { return a[i].index < a[j].index }

package doco

import "io/ioutil"

// Document represents a documentation file
// that can be parsed and converted to HTML.
type Document interface {
	Path() string
	ReadContent() ([]byte, error)
}

type Documents []Document

func (d Documents) Meta() Document {
	for _, v := range d {
		if v.Path() != "docs/_meta.md" {
			continue
		}

		return v
	}

	return nil
}

// fileDocument is a Document implementation
// based on the path to the file. It assumes the file
// will be readable by the current user.
type fileDocument struct {
	path string
}

// Path of the file document.
func (f fileDocument) Path() string {
	return f.path
}

// ReadContent of the file document by reading the file.
func (f fileDocument) ReadContent() ([]byte, error) {
	return ioutil.ReadFile(f.path)
}

// NewFile returns a new Document implementation
// based on the path to the file.
func NewFile(path string) Document {
	return fileDocument{path: path}
}

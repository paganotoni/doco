package doco

// Site being parsed and built from the
// docs markdown files.
type Site struct {
	// List of documents parsed from the docs folder
	// files starting with underscore are not parsed
	documents Documents
}

// add a document to the site. Useful when walking
// a directory to add the files to the site.
func (s *Site) add(d Document) {
	s.documents = append(s.documents, d)
}

func (s *Site) Documents() Documents {
	return s.documents
}

func NewSite(docs Documents) *Site {
	site := &Site{
		documents: docs,
	}

	return site
}

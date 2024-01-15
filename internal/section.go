package internal

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

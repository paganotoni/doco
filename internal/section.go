package internal

type section struct {
	name string
	path string

	documents []document
}

func (s *section) String() string {
	pp := s.name + "\n"
	for _, doc := range s.documents {
		pp += "      " + doc.String() + "\n"
	}

	return pp
}

type sections []section

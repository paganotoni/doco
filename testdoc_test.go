package doco_test

// testdocument is a Document implementation
// based on a slice where the first element
// is the path and the second is the content.
type testdocument []string

// Path of the test document.
func (tf testdocument) Path() string {
	return tf[0]
}

// ReadContent of the test document.
func (tf testdocument) ReadContent() ([]byte, error) {
	return []byte(tf[1]), nil
}

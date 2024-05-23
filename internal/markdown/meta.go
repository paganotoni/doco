package markdown

import (
	"bytes"

	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// ReadMetadata parses the content and returns the metadata
// from the markdown file.
func ReadMetadata(content []byte) (map[string]interface{}, error) {
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := mparser.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	return meta.Get(context), nil
}

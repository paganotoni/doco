package markdown

import (
	"bytes"
	"html/template"
)

// HTMLFrom the passed content.
func HTMLFrom(content []byte) (template.HTML, error) {
	var buf bytes.Buffer
	if err := mparser.Convert(content, &buf); err != nil {
		return template.HTML(""), err
	}

	return template.HTML(buf.String()), nil
}

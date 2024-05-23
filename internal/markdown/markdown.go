// main package takes care of the Markdown operations
// for Doco. It is useful to read metadata from the markdown
// files and convert the markdown content to HTML.
package markdown

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
)

// mparser is the markdown parser used to parse the content
var mparser = goldmark.New(
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(), // read note
	),

	goldmark.WithExtensions(
		meta.Meta,

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

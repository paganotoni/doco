package web

import (
	"testing"
)

func TestNavigation(t *testing.T) {
	t.Run("Basic Case", func(t *testing.T) {
		pages := Pages{
			{SourcePath: "docs/_meta.md", Metadata: map[string]interface{}{}},
			{SourcePath: "docs/index.md", Metadata: map[string]interface{}{"Title": "Home"}},
			{SourcePath: "docs/getting-started.md", Metadata: map[string]interface{}{"Title": "Getting Started"}},
		}

		nn := navigation(pages).HTMLFor(pages[0])
		expected := `<ul><li><a href="index.html">Home</a></li><li><a href="getting-started.html">Getting Started</a></li></ul>`
		if string(nn) != expected {
			t.Fatalf("expected to be '%s' but got '%s'", expected, nn)
		}
	})

	t.Run("Weight involved", func(t *testing.T) {
		pages := Pages{
			{SourcePath: "docs/getting-started.md", Metadata: map[string]interface{}{"Title": "Getting Started", "Weight": 2}},
			{SourcePath: "docs/index.md", Metadata: map[string]interface{}{"Title": "Home", "Weight": 1}},
			{SourcePath: "docs/other.md", Metadata: map[string]interface{}{"Title": "Other", "Weight": 3}},
		}

		nn := navigation(pages).HTMLFor(pages[0])
		expected := `<ul><li><a href="index.html">Home</a></li><li><a href="getting-started.html">Getting Started</a></li><li><a href="other.html">Other</a></li></ul>`
		if string(nn) != expected {
			t.Fatalf("expected to be '%s' but got '%s'", expected, nn)
		}
	})
}

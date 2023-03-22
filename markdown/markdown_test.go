package markdown_test

import (
	"strings"
	"testing"

	"github.com/paganotoni/doco/markdown"
)

func TestMetadataFrom(t *testing.T) {
	source := []byte(`---
title: "My Title"
description: "My Description"
---`)

	meta, err := markdown.MetadataFrom(source)
	if err != nil {
		t.Fatalf("error parsing metadata: %s", err)
	}

	if meta["title"] != "My Title" {
		t.Fatalf("expected title to be 'My Title' but got '%s'", meta["title"])
	}

	if meta["Title"] != nil {
		t.Fatalf("expected title to be nil but got '%s'", meta["Title"])
	}
}

func TestHTMLFrom(t *testing.T) {
	source := []byte(`---
title: "My Title"
description: "My Description"
---
# My Title
`)

	html, err := markdown.HTMLFrom(source)
	if err != nil {
		t.Fatalf("error parsing html: %s", err)
	}

	expected := `<h1>My Title</h1>`
	if strings.TrimSpace(string(html)) != expected {
		t.Fatalf("expected html to be '%s' but got '%s'", expected, html)
	}
}

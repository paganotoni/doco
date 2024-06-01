package internal

import "testing"

func TestMeta(t *testing.T) {
	var source = `---
Title: something
Summary: Markdown document
Tags:
   - markdown
   - goldmark
---

# Hello custom meta parsing
`
	meta, err := parseMeta([]byte(source))
	if err != nil {
		t.Fatalf("Error parsing meta: %s", err)
	}

	if meta["Title"] != "something" {
		t.Fatalf("Expected Title to be 'something', got %s", meta["Title"])
	}
	a, ok := meta["Tags"].([]interface{})
	if !ok {
		t.Fatalf("Expected Tags to be a slice of strings, got %T", meta["Tags"])
	}

	if len(a) != 2 {
		t.Fatalf("Expected Tags to have 2 elements, got %d", len(a))
	}
}

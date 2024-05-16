package main

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	bb := bytes.NewBuffer([]byte{})
	if err := printHelp(bb); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(bb.String(), help) {
		t.Fatal("expected output")
	}

	if !strings.Contains(bb.String(), "--folder") {
		t.Fatal("expected flags")
	}

	if !strings.Contains(bb.String(), "--output") {
		t.Fatal("expected flags")
	}
}

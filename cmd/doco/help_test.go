package main

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
)

//go:embed help.txt
var content string

func TestRun(t *testing.T) {
	bb := bytes.NewBuffer([]byte{})
	if err := printHelp(bb); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(bb.String(), content) {
		t.Fatal("expected output")
	}

	if !strings.Contains(bb.String(), "--folder") {
		t.Fatal("expected flags")
	}

	if !strings.Contains(bb.String(), "--output") {
		t.Fatal("expected flags")
	}
}

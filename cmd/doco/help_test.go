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
}

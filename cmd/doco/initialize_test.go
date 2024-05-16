package main

import (
	"bytes"
	"os"
	"path"
	"testing"
)

func TestInitialize(t *testing.T) {
	f := path.Join(t.TempDir(), "docs")

	err := initialize(f, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	// Test if the directory was created
	if _, err := os.Stat(f); err != nil {
		t.Fatal(err)
	}

	files := []string{
		"index.md",
		"getting_started.md",
		"assets/logo.png",
		"assets/favicon.png",
		"assets/preview.png",
		"_meta.md",
	}

	for _, v := range files {
		if _, err := os.Stat(path.Join(f, v)); err != nil {
			t.Fatal(err)
		}
	}

}

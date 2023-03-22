package doco

import (
	"fmt"
	"os"
	"path/filepath"
)

// docsFolder to scan for documentation files.
const docsFolder = "docs"

// Parse the files in the docs folder and return a Site
// containing those files. The web package may use returned
// Site instance to generate HTML for the documentation.
func Parse() (site *Site, err error) {
	site = &Site{}
	err = filepath.Walk(docsFolder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		site.add(NewFile(path))
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find files: %v", err)
	}

	return site, err
}

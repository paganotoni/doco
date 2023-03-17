package doco

import (
	"fmt"
	"os"
	"path/filepath"
)

func Parse() (docs Documents, err error) {
	err = filepath.Walk("docs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		docs = append(docs, Document(path))
		return nil
	})

	if err != nil {
		return docs, fmt.Errorf("failed to find files: %v", err)
	}

	return docs, nil
}

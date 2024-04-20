package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Generates the initial documentation structure
func initialize(folder string) error {
	fmt.Println("> Initializing docs folder")
	_, err := os.Stat(folder)
	if err == nil {
		return errors.New("folder exists, aborting init")
	}

	// Create folder
	err = os.Mkdir(folder, 0755)
	if err != nil {
		return fmt.Errorf("error creating the folder: err")
	}

	// Create assets folder
	err = os.Mkdir(filepath.Join(folder, "assets"), 0755)
	if err != nil {
		return fmt.Errorf("error creating the folder: err")
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
		f, err := base.Open("base/" + v)
		if err != nil {
			return fmt.Errorf("error opening file %s: %w", v, err)
		}
		defer f.Close()

		// Create file
		file, err := os.Create(folder + "/" + v)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", v, err)
		}
		defer file.Close()

		// Copy content
		_, err = io.Copy(file, f)
		if err != nil {
			return fmt.Errorf("error copying file %s: %w", v, err)
		}
	}

	return nil
}

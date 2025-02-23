package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var ignoreFolders = []string{
	"node_modules",
	".git",
}

// watch a folder for changes and rebuild the documentation
// when a change is detected.
func watch(docsFolder, dstFolder string) {
	log.Println("> 👀 Watching for changes in", docsFolder)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Ignoring changes in destination folder.
				if isFileInDirectory(event.Name, dstFolder) {
					return
				}

				build(docsFolder, dstFolder)
				log.Println(">", event.Name, "changed, documentation rebuilt.")

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// add the folder and its subfolders to the watcher by
	// walking the folder.
	err = filepath.Walk(docsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isInFolder(".git", path) {
			return nil
		}

		if isInFolder("node_modules", path) {
			return nil
		}

		err = watcher.Add(path)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

// Serve the public folder
func serve(dstFolder string) error {
	fs := http.FileServer(http.Dir(dstFolder))
	http.Handle("/", fs)

	log.Println("> Serving documentation on http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}

	return nil
}

// check if a file path is within a directory
func isFileInDirectory(filePath, directory string) bool {
	// Convert both paths to absolute paths
	absFile, err := filepath.Abs(filePath)
	if err != nil {
		return false
	}
	absDir, err := filepath.Abs(directory)
	if err != nil {
		return false
	}

	// Use filepath.HasPrefix to check if the file is in the directory
	return strings.HasPrefix(absFile, absDir)
}

// Check if a file path contains .git in its path components
func isInFolder(folder, path string) bool {
	// Split path into components
	components := strings.Split(filepath.Clean(path), string(filepath.Separator))

	// Check if any component is .git
	for _, comp := range components {
		if comp == folder {
			return true
		}
	}
	return false
}

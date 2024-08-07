package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

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

package main

import (
	"log"
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

				if event.Has(fsnotify.Write) {
					build(docsFolder, dstFolder)
					log.Println(">", event.Name, "changed, documentation rebuilt.")
				}
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

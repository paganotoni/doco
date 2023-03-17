package doco

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
)

func Serve() error {
	ch := make(chan bool)
	go Build()
	go watch(ch, "docs")

	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("[INFO] Serving docs on http://localhost:8080 (Press Ctrl+C to stop)")
	err := http.ListenAndServe(":8080", nil)
	ch <- true

	return err
}

func watch(ch chan bool, folder string) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	// Start listening for events.
	go func() {
		fmt.Println("[INFO] Watching for changes")

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					fmt.Println("[INFO] Changes detected")
					Build()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(folder)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-ch
	fmt.Println("Done")
}

package serve

import (
	"log"
	"net/http"
)

// Serve the public folder
func Run() error {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Println("> Serving documentation on http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}

	return nil
}

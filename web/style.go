package web

import (
	_ "embed"

	"fmt"
	"io/ioutil"
)

var (
	//go:embed templates/styles.css
	style []byte
)

// Writes the style.css file to the public folder. This style has been
// precompiled with tailwindCSS.
func writeStyles() error {
	fmt.Println("[INFO] Write > public/style.css")
	err := ioutil.WriteFile("public/style.css", style, 0777)
	if err != nil {
		return err
	}

	return nil
}

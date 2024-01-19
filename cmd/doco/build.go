package main

import (
	"errors"
	"os"

	"github.com/paganotoni/doco/internal"
)

// build generates the static html files for the site
// and writes them to the destination folder.
func build(src, dst string) error {
	_, err := os.Stat(src)
	if err != nil {
		return errors.New("docs folder does not exist, aborting build")
	}

	site, err := internal.NewSite(src)
	if err != nil {
		return err
	}

	return internal.Generate(src, dst, site)
}

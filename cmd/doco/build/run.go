package build

import (
	"os"

	"github.com/paganotoni/doco"
	"github.com/paganotoni/doco/web"
)

// buildFolder is the folder where the resulting files
// will be written to
const buildFolder = "public"

func Run() error {
	// Cleaning the build folder
	if err := os.RemoveAll(buildFolder); err != nil {
		return err
	}

	if err := os.MkdirAll(buildFolder, 0777); err != nil {
		return err
	}

	site, err := doco.Parse()
	if err != nil {
		return err
	}

	err = web.Generate(site)
	if err != nil {
		return err
	}

	// docs, err := doco.Parse()
	// if err != nil {
	// 	return err
	// }

	return nil
}

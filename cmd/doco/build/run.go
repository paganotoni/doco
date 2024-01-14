package build

import "github.com/paganotoni/doco/internal"

// dstFolder is the folder where the resulting files
// will be written to
var dstFolder = "public"

// srcFolder is the folder where the source files
// are located
var srcFolder = "docs"

// Run generates the static html files for the site
// and writes them to the destination folder.
func Run() error {
	site, err := internal.NewSite(srcFolder)
	if err != nil {
		return err
	}

	return internal.Generate(srcFolder, dstFolder, site)
}

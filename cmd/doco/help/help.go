package help

import (
	_ "embed"
	"fmt"

	"io"
)

//go:embed help.txt
var help string

// Run the help command
func Run(w io.Writer) error {
	fmt.Fprintf(w, help)

	return nil
}

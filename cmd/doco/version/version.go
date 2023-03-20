package version

import (
	"fmt"
	"io"

	"github.com/paganotoni/doco"
)

// Runs the version command
func Run(w io.Writer) error {
	fmt.Fprintf(w, "Doco %v", doco.Version)

	return nil
}

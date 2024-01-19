package main

import (
	_ "embed"
	"flag"
	"fmt"

	"io"
)

//go:embed help.txt
var help string

// Run the help command
func printHelp(w io.Writer) error {
	fmt.Fprintf(w, help)
	fmt.Fprintf(w, "\nOptions:\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(w, "  --"+f.Name+"\t"+f.Usage+"\n")
	})

	return nil
}

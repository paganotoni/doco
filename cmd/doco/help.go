package main

import (
	_ "embed"
	"flag"
	"fmt"

	"io"
)

const help = `Doco is a simple, lightweight, and easy to use documentation generator for Markdown files. It takes a directory of Markdown files and generates a static site with a table of contents.

Usage
    doco command [arguments]

Commands
    build     Generates the documentation site into the /public directory (see options)
    help      Prints the CLI help
    init      Creates a /docs directory with base files (see options)
    serve     Starts a local server to view the documentation site
`

// Run the help command
func printHelp(w io.Writer) error {
	fmt.Fprintf(w, help)
	fmt.Fprintf(w, "\nOptions:\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(w, "  --"+f.Name+"\t"+f.Usage+"\n")
	})

	return nil
}

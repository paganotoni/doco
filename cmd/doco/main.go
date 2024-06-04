package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// docs folder is the folder where the documentation
	// is stored and updated.
	docsFolder string

	// dstFolder is the folder where the resulting files
	// will be written to
	dstFolder string
)

func init() {
	flag.StringVar(&docsFolder, "folder", "docs", "source folder for the documentation")
	flag.StringVar(&dstFolder, "output", "public", "folder to put the generated files")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("No command specified %s\n", args)
		fmt.Println("--------")
		printHelp(os.Stdout)

		return
	}

	switch args[0] {
	case "build":
		err := build(docsFolder, dstFolder)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

	case "init":
		err := initialize(docsFolder, os.Stdout)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("Initialized docs folder")
	case "serve":
		build(docsFolder, dstFolder)

		go watch(docsFolder, dstFolder)
		serve(dstFolder)
	case "help":
		printHelp(os.Stdout)
	default:
		fmt.Printf("Unknown command %s\n", args[0])
		fmt.Println("--------")
		printHelp(os.Stdout)
	}

	return
}

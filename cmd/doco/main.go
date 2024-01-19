package main

import (
	"embed"
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
	dstFolder = "public"

	// the base folder contains the initial files
	// that are copied to the docs folder when init rans.
	//
	//go:embed all:base
	base embed.FS
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

		// TODO: Run build by default ?
		return
	}

	switch args[0] {
	case "build":
		err := build(docsFolder, dstFolder)
		if err != nil {
			fmt.Println(err)
		}

	case "init":
		err := initialize(docsFolder)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("Initialized docs folder")
	case "help":
		printHelp(os.Stdout)
	case "serve":
		serve()
	default:
		fmt.Printf("Unknown command %s\n", args[1])
		fmt.Println("--------")
		printHelp(os.Stdout)
	}

	return
}

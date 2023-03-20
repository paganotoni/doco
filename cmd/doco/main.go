package main

import (
	"fmt"
	"os"

	"github.com/paganotoni/doco/cmd/doco/build"
	"github.com/paganotoni/doco/cmd/doco/help"
	"github.com/paganotoni/doco/cmd/doco/version"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		// err := doco.Build()
		// if err != nil {
		// 	fmt.Println(err)
		// }

		return
	}

	switch args[1] {
	case "build":
		err := build.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "serve":
		// doco.Serve()
	case "help":
		help.Run(os.Stdout)
	case "version":
		version.Run(os.Stdout)
	default:
		fmt.Printf("Unknown command %s\n", args[1])
		fmt.Println("--------")
		help.Run(os.Stdout)
	}

	return
}

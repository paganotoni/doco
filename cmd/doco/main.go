package main

import (
	"fmt"
	"os"

	"github.com/paganotoni/doco/cmd/doco/build"
	"github.com/paganotoni/doco/cmd/doco/help"
	"github.com/paganotoni/doco/cmd/doco/serve"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		// Run build by default ?
		return
	}

	switch args[1] {
	case "build":
		err := build.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "help":
		help.Run(os.Stdout)
	case "serve":
		serve.Run()
	default:
		fmt.Printf("Unknown command %s\n", args[1])
		fmt.Println("--------")
		help.Run(os.Stdout)
	}

	return
}

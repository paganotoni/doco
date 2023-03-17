package main

import (
	"fmt"
	"os"

	"github.com/paganotoni/doco"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		err := doco.Build()
		if err != nil {
			panic(err)
		}

		return
	}

	switch args[1] {
	case "serve":
		doco.Serve()
	case "help":
		fmt.Println("help")
	case "version":
		fmt.Println("v0.0.1")
	default:
		fmt.Printf("Unknown command %s\n", args[1])
	}

	return
}

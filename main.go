package main

import (
	"fmt"
	"os"

	"github.com/hardfist/gopack/pkg/cli"
)

func main() {
	app := cli.New()
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

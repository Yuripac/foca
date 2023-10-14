package main

import (
	"fmt"
	"os"

	"github.com/yuripac/foca/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		fmt.Println("Something went wrong...")

		os.Exit(1)
	}
}


package main

import (
	"log"
	"os"

	"github.com/yuripac/foca/cmd/foca"
)

func main() {
	if err := foca.Command().Execute(); err != nil {
		log.Println("Something went wrong...")

		os.Exit(1)
	}
}

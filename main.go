package main

import (
	"log"
	"os"
)

func main() {
	if err := Command().Execute(); err != nil {
		log.Println("Something went wrong...")

		os.Exit(1)
	}
}

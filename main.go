package main

import (
	"github.com/yuripac/foca/schedule"
)

func main() {
	err := schedule.Init()
	if err != nil {
		panic(err)
	}

	s, err := schedule.Load()
	if err != nil {
		panic(err)
	}

	s.Run()
}

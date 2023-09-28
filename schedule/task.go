package schedule

import (
	"log"
	"sync"
)

type Task struct {
	Title    string
	Cron     string
	Desc     string    `yaml:"description"`
	Commands []Command `yaml:"commands"`
}

func (t Task) Run() {
	wg := sync.WaitGroup{}

	for _, c := range t.Commands {
		wg.Add(1)
		go func(c Command) {
			if err := c.Run(); err != nil {
				log.Printf("error running the command: \"%s\": %s\n", c.TODO, err)
			} else {
				log.Printf("command \"%s\" finished successfuly!\n", c.Name())
			}
			wg.Done()
		}(c)
	}
	wg.Wait()
}

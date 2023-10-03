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

	RunWg sync.WaitGroup
}

func (t *Task) Run() {
	for _, c := range t.Commands {
		t.RunWg.Add(1)
		go func(c Command) {
			if err := c.Run(); err != nil {
				log.Printf("error running the command: \"%s\": %s\n", c.TODO, err)
			} else {
				log.Printf("command \"%s\" finished successfuly!\n", c)
			}
			t.RunWg.Done()
		}(c)
	}
}

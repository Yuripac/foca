package schedule

import (
	"fmt"
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
				fmt.Printf("error running the command: \"%s\": %s\n", c.TODO, err)
			} else {
				fmt.Printf("command \"%s\" finished successfuly!\n", c)
			}
			t.RunWg.Done()
		}(c)
	}
}

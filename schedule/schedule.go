package schedule

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v3"
)

type Schedule struct {
	Tasks map[string]*Task

	RunWg sync.WaitGroup
}

const configPath = "config/schedule_example.yaml"

func (s *Schedule) Run() {
	for _, t := range s.Tasks {
		s.RunWg.Add(1)
		go func(t *Task) {
			t.Run()
			t.RunWg.Wait()

			s.RunWg.Done()
		}(t)
	}
}

func (s *Schedule) Cron() *cron.Cron {
	c := cron.New()
	for name, t := range s.Tasks {
		t := t
		c.AddFunc(t.Cron, func() {
			// t.Run()
			// t.RunWg.Wait()

			home, _ := os.UserHomeDir()
			cmd := exec.Command(home+"/go/bin/foca", "terminal-exec-task", name)
			if err := cmd.Run(); err != nil {
				log.Println("error on terminal-exec-task:", err)
			}
		})
	}
	return c
}

func Init(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("file \"%s\" already exists", path)
	}

	dir := strings.Split(path, "/")
	dir = dir[0:(len(dir) - 1)]
	if err := os.MkdirAll(strings.Join(dir, "/"), 0777); err != nil {
		return err
	}

	text, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, text, 0777); err != nil {
		return err
	}

	return nil
}

func Load(path string) (*Schedule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sch := Schedule{}
	if err = yaml.NewDecoder(f).Decode(&sch.Tasks); err != nil {
		return nil, err
	}

	return &sch, nil
}

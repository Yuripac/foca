package schedule

import (
	"fmt"
	"os"
	"strings"
	"sync"

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

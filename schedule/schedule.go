package schedule

import (
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

const initialConfig = `study:
  title: Study focus!!!
  cron: 0 18 * * *
  commands:
  - title: notion
    command: xdg-open https://notion.so
  - command: code ~/my_project
    workspace: 2
  `

type Schedule struct {
	Tasks map[string]*Task

	RunWg sync.WaitGroup
}

var (
	home, _ = os.UserHomeDir()
	path    = home + "/.foca/schedule.yaml"
)

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

func Init() error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dir := strings.Split(path, "/")
		dir = dir[0:(len(dir) - 1)]

		if err := os.MkdirAll(strings.Join(dir, "/"), 0777); err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write([]byte(initialConfig)); err != nil {
			return err
		}
	}

	return nil
}

func Load() (*Schedule, error) {
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

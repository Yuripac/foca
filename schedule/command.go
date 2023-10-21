package schedule
 
import (
	"os/exec"
	"strings"
)

type Command struct {
	Title	  string
	TODO      string `yaml:"command"`
	Workspace int
}

func (c Command) Run() error {
	todo := strings.Split(c.TODO, " ")
	name, args := todo[0], todo[1:]

	err := exec.Command(name, args...).Run()
	if err != nil {
		return err
	}

	return nil
}

func (c Command) String() string {
	if c.Title != "" {
		return c.Title
	}

	return c.TODO
}

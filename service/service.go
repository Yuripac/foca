package service

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

type Service struct {
	User, ExecStart, Path string
}

const configPath = "config/foca.service"

func (s *Service) Import() error {
	if s.Path == "" {
		return fmt.Errorf("path is missing")
	}

	dir := strings.Split(s.Path, "/")
	dir = dir[0:(len(dir) - 1)]
	if err := os.MkdirAll(strings.Join(dir, "/"), 0755); err != nil {
		return err
	}

	t, err := template.New("foca.service").ParseFiles(configPath)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(s.Path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	return t.Execute(f, s)
}

func (s *Service) Start() error {
	cmd := exec.Command("systemctl", "--user", "daemon-reload")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("systemctl", "--user", "restart", s.Filename())
	return cmd.Run()
}

func (s Service) Stop() error {
	cmd := exec.Command("systemctl", "--user", "stop", s.Filename())

	return cmd.Run()
}

func (s *Service) Filename() string {
	split := strings.Split(s.Path, "/")
	return split[len(split)-1]
}

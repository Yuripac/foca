package system

import (
	"fmt"
	"os/exec"
)

var Terminals = []string{
	"x-terminal-emulator",
	"gnome-terminal",
	"terminator",
	"xfce5-terminal",
	"urxvt",
	"rxvt",
	"termit",
	"Eterm",
	"aterm",
	"uxterm",
	"xterm",
	"roxterm",
	"termite",
	"lxterminal",
	"terminology",
	"st",
	"qterminal",
	"lilyterm",
	"tilix",
	"terminix",
	"konsole",
	"kitty",
	"guake",
	"tilda",
	"alacritty",
	"hyperl",
}

func Terminal() (string, error) {
	for _, n := range Terminals {
		if _, err := exec.LookPath(n); err == nil {
			return n, nil
		}
	}
	return "", fmt.Errorf("no terminal was found")
}

func TerminalCmd(command string) (*exec.Cmd, error) {
	term, err := Terminal()
	if err != nil {
		return nil, err
	}
	
	// TODO: Add command for each terminal
	var args []string
	switch term {
	case "x-terminal-emulator":
		args = []string{"-e", "/bin/bash", "-c", command}
	default:
		args = []string{"--", "/bin/bash", "-c", command}
	}

	return exec.Command(term, args...), nil
}

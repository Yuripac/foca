package cmd

import (
	"strings"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func TerminalExecTaskCmd() *cobra.Command {
	cmd := cobra.Command {
		Use: "terminal-exec-task",
		Short: "Open a terminal for task execution confirmation",
		Run: func(_ *cobra.Command, args []string) {
			execTask := fmt.Sprintf(
			u.HomeDir+"/go/bin/foca exec-task %s; sleep inifity",
			strings.Join(args, " "),
	)

			cmd := exec.Command("gnome-terminal", "--", "/bin/bash", "-c", execTask)
			if err := cmd.Run(); err != nil {
				fmt.Println("Error ao inicializar o terminal: ", err)
			}

		},
	}

	return &cmd
}

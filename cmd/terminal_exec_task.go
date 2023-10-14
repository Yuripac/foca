package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuripac/foca/system"
)

func TerminalExecTaskCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "terminal-exec-task",
		Short: "Open a terminal for task execution confirmation",
		Run: func(_ *cobra.Command, args []string) {
			execTask := fmt.Sprintf(
				u.HomeDir+"/go/bin/foca exec-task %s; sleep inifity",
				strings.Join(args, " "),
			)

			cmd, err := system.TerminalCmd(execTask)
			if err != nil {
				fmt.Println(err)

				os.Exit(1)
			}

			if err := cmd.Run(); err != nil {
				fmt.Println("Error ao inicializar o terminal: ", err)
			}
		},
	}

	return &cmd
}

package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
	"github.com/yuripac/foca/system"
)

var (
	u, _    = user.Current()
	schPath = u.HomeDir + "/.config/foca/schedule.yaml"

	s = system.Service{
		User:      u.Username,
		ExecStart: u.HomeDir + "/go/bin/foca start-cron",
		Path:      u.HomeDir + "/.config/systemd/user/foca.service",
	}
)

func RootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "foca",
		Short: "Make easy scheduling your tasks",
		Run:   func(_ *cobra.Command, _ []string) {},
	}

	rootCmd.AddCommand(InitCmd())
	rootCmd.AddCommand(StartCronCmd())
	rootCmd.AddCommand(StartServiceCmd())
	rootCmd.AddCommand(StopServiceCmd())
	rootCmd.AddCommand(ExecTaskCmd())
	rootCmd.AddCommand(TerminalExecTaskCmd())

	return &rootCmd
}

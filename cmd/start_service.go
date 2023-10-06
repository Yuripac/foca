package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

func StartServiceCmd() *cobra.Command {
	c := cobra.Command {
		Use: "start-service",
		Short: "Start a service to execute the schedule",
		Run: func(_ *cobra.Command, _ []string) {
			if err := s.Import(); err != nil {
				fmt.Println("error on start-service:", err)

				os.Exit(1)
			}

			if err := s.Start(); err != nil {
				fmt.Println("error on start-service:", err)

				os.Exit(1)
			}
		},
	}

	return &c
}

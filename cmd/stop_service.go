package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

func StopServiceCmd() *cobra.Command {
	c := cobra.Command {
		Use: "stop-service",
		Short: "",
		Run: func(_ *cobra.Command, _ []string) {
			if err := s.Stop(); err != nil {
				fmt.Println("error on stop-service:", err)

				os.Exit(1)
			}
		},
	}

	return &c
}

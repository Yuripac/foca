package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/yuripac/foca/schedule"
)

func InitCmd() *cobra.Command {
	c := cobra.Command {
		Use: "init",
		Short: "",
		Run: func(_ *cobra.Command, _ []string){
			if err := schedule.Init(schPath); err != nil && !os.IsExist(err) {
				fmt.Println("error on init:", err)

				os.Exit(1)
			}

			editCmd := exec.Command("vi", schPath)
			editCmd.Stdin, editCmd.Stdout, editCmd.Stderr = os.Stdin, os.Stdout, os.Stderr

			if err := editCmd.Run(); err != nil {
				fmt.Println(err)
			}
		},
	}

	return &c
}

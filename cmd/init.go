package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuripac/foca/schedule"
	"fmt"
	"os"
)

func InitCmd() *cobra.Command {
	c := cobra.Command {
		Use: "init",
		Short: "",
		Run: func(_ *cobra.Command, _ []string){
			
			if err := schedule.Init(schPath); err != nil {
				fmt.Println("error on init:", err)

				os.Exit(1)
			}
		},
	}

	return &c
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuripac/foca/schedule"
)

func ExecTaskCmd() *cobra.Command {
	var confirm bool
	c := cobra.Command{
		Use:   "exec-task task",
		Short: "Stop the service",
		Run: func(_ *cobra.Command, args []string) {
			s, err := schedule.Load(schPath)
			if err != nil {
				fmt.Println("error loading schedule:", err)

				os.Exit(1)
			}

			for _, name := range args {
				t := s.Tasks[name]
				if !confirm || askConfirm(name) {
					t.Run()
					t.RunWg.Wait()
				}
			}
		},
	}
	c.Flags().BoolVar(&confirm, "confirm", true, "Pede confirmação")

	return &c
}

func askConfirm(t string) bool {
	fmt.Printf("The task \"%s\" will be executed, do you want to continue? [Y/n]: ", t)
	var input string

	if _, err := fmt.Scanln(&input); err != nil {
		input = "Y"
	}
	
	return strings.ToUpper(input) == "Y"
}

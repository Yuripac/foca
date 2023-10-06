package cmd

import (
	"fmt"
	"os"
	"slices"

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

			for _, a := range args {
				t := s.Tasks[a]
				if confirm {
					fmt.Printf("A task \"%s\" será executada a seguir, deseja continuar? [Y/n]: ", a)
					var input string
					if _, err := fmt.Scanln(&input); err != nil {
						input = "Y"
					}

					if slices.Contains([]string{"y", "Y"}, input) {
						t.Run()
						t.RunWg.Wait()
					}
				} else {
					t.Run()
					t.RunWg.Wait()
				}

			}
		},
	}
	c.Flags().BoolVar(&confirm, "confirm", true, "Pede confirmação")

	return &c
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuripac/foca/schedule"
	"fmt"
	"os"
	"context"
)

func StartCronCmd() *cobra.Command {
	c := cobra.Command {
		Use: "start-cron",
		Short: "Start to execute the schedule",
		Run: func(_ *cobra.Command, _ []string) {
			ctx := context.Background()

			s, err := schedule.Load(schPath)
			if err != nil {
				fmt.Println("error loading schedule:", err)

				os.Exit(1)
			}

			go func(s *schedule.Schedule) { s.Cron().Run() }(s)

			<-ctx.Done()
		},
	}

	return &c
}

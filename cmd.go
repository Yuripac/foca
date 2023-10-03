package main

import (
	"context"
	"fmt"
	"os"
	"os/user"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/yuripac/foca/schedule"
	"github.com/yuripac/foca/service"
)

var (
	u, _    = user.Current()
	schPath = u.HomeDir + "/.foca/schedule.yaml"
)

func Command() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "foca",
		Short: "Make easy schedule your tasks",
		Run:   func(_ *cobra.Command, _ []string) {},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Create a initial schedule to you edit",
		Run:   Init,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start-cron",
		Short: "Start to execute the schedule",
		Run:   StartCron,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start-service",
		Short: "Start a service to execute the schedule",
		Run:   StartService,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "stop-service",
		Short: "Stop the service",
		// Run:   StopService,
	})

	return &rootCmd
}

func Init(cmd *cobra.Command, args []string) {
	if err := schedule.Init(schPath); err != nil {
		fmt.Println("error on init:", err)

		os.Exit(1)
	}
}

func StartCron(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	s, err := schedule.Load(schPath)
	if err != nil {
		fmt.Println("error loading schedule:", err)

		os.Exit(1)
	}

	go func(s *schedule.Schedule) {
		c := cron.New()
		for _, t := range s.Tasks {
			t := t
			c.AddFunc(t.Cron, func() {
				t.Run()
				t.RunWg.Wait()
			})
		}
		c.Start()
	}(s)

	<-ctx.Done()
}

func StartService(cmd *cobra.Command, args []string) {
	s := service.Service{
		User:      u.Username,
		ExecStart: u.HomeDir + "/go/bin/foca start-cron",
		Path:      u.HomeDir + "/.config/systemd/user/foca.service",
	}

	if err := s.Import(); err != nil {
		fmt.Println("error on start-service:", err)

		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println("error on start-service:", err)

		os.Exit(1)
	}
}

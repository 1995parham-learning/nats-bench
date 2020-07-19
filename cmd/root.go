package cmd

import (
	"os"

	"github.com/1995parham/nats-bench/cmd/producer"
	"github.com/1995parham/nats-bench/cmd/subscriber"
	"github.com/1995parham/nats-bench/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := &cobra.Command{
		Use:   "nats-bench",
		Short: "NATS Benchmark tools and guidelines",
	}

	config.Register(root.PersistentFlags())

	producer.Register(root)
	subscriber.Register(root)

	if err := root.Execute(); err != nil {
		logrus.Errorf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}

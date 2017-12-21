package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCommand = cobra.Command{
		Use:   "xamanek",
		Short: "xamanek - Cloudflavor FaaS node scheduler",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: this should be the entry point of the app.
		},
	}

	configPath string
)

func main() {
	if err := rootCommand.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCommand.PersistentFlags().StringVar(
		&configPath,
		"config",
		"",
		"Full path to the config file",
	)
	logrus.SetLevel(logrus.DebugLevel)
}

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudflavor/xamanek/pkg/core"
)

var (
	configPath string

	rootCommand = cobra.Command{
		Use:   "xamanek",
		Short: "xamanEk - Cloudflavor FaaS node scheduler",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: this should be the entry point of the app.
			viper.AddConfigPath(".")
			if configPath != "" {
				viper.AddConfigPath(configPath)
			}
			if err := viper.ReadInConfig(); err != nil {
				logrus.Fatalf("Failed to read config: %v", err)
			}
			newCfg := core.NewConfig()
			if err := viper.Unmarshal(newCfg); err != nil {
				logrus.Fatalf("Error while unmarshalling configuration: %v", err)
			}
			instance := core.NewInstance(newCfg)
			if err := instance.Start(); err != nil {
				logrus.Fatalf("Error while starting application services: %v", err)
			}
		},
	}
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

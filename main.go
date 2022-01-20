package main

import (
	"os"

	"github.com/omegion/argocd-actions/cmd"
	cmd2 "github.com/omegion/cobra-commander"
)

const (
	// Config file name where a config file will be created.
	// For example, $HOME/.argocdActions/config.yaml.
	configFileName = "argocdActions"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --token is bound to ACDA_TOKEN.
	configEnvPrefix = "ARGOCD"
)

func main() {
	commander := cmd2.NewCommander(cmd.Root()).
		SetCommand(
			cmd.Version(),
			cmd.Sync(),
		).
		SetConfig(getConfig()).
		Init()

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}

func getConfig() *cmd2.Config {
	configName := configFileName
	environmentPrefix := configEnvPrefix

	return &cmd2.Config{Name: &configName, EnvironmentPrefix: &environmentPrefix}
}

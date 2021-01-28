package cmd

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: FormattedMessage(),
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in Server file, command line or ENV variables if set
func initConfig() {
	viper.SetEnvPrefix("SRV")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Panic(err)
		os.Exit(1)
	}
}

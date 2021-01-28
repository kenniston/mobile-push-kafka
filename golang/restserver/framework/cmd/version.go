package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//===============================================================================
// Default build-time variable
// These values are overridden via ldflags
var (
	Version   = "0.1"
	GitCommit = "-"
	BuildTime = "-"
)

const versionF = `Server
  Version: %s
  GitCommit: %s
  BuildTime: %s

`

// FormattedMessage get full formatted version message
func FormattedMessage() string {
	return fmt.Sprintf(versionF, Version, GitCommit, BuildTime)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Microservice build version, build time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(FormattedMessage())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
"fmt"
"github.com/spf13/cobra"
)

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version   = "0.1"
	GitCommit = "unknown-commit"
	BuildTime = "unknown-buildtime"
)

const versionF = `mobilepushproducer
  Version: %s
  GitCommit: %s
  BuildTime: %s
`

// FormattedMessage gets the full formatted version message
func FormattedMessage() string {
	return fmt.Sprintf(versionF, Version, GitCommit, BuildTime)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display this build's version, build time, and git hash",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(FormattedMessage())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
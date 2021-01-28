package cmd

import (
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Starts the Server",
	Long: `This Server provides common endpoints
for integration with other platforms and systems.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := framework.NewBasicServer(viper.GetViper())

		if err := server.Run(); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

func GetRunCommand() *cobra.Command {
	return runCmd
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Configure Server Commands
	runCmd.Flags().StringP("server-port", "p", "4001", "Configure Server port")
	runCmd.Flags().StringP("log-level", "l", "info", "Configure log level")

	err := viper.GetViper().BindPFlags(runCmd.Flags())
	if err != nil {
		panic(err)
	}
}

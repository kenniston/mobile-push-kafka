package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command {
	Use: "run",
	Aliases: []string{"r"},
	Short: "Starts the Kafka Producer REST Microservice",
	Long: `The Kafka Producer REST Microservice provides endpoints to send push 
message to Kafka server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("port", "p", viper.GetString("port"), "Configure microservice port")

	err := viper.GetViper().BindPFlags(runCmd.Flags())
	if err != nil {
		panic(err)
	}
}
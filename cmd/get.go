package cmd

import (
	"d3u/bboltconnection"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List games",
	Long:  `List all games with their location info`,
	Run: func(cmd *cobra.Command, args []string) {
		bboltconnection.GetGames()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

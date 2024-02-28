package cmd

import (
	"d3u/bboltconnection"
	"github.com/spf13/cobra"
)

// rmCmd represents the remove command
var rmCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove game data",
	Long:  `Remove game's data from db'`,
	Run: func(cmd *cobra.Command, args []string) {

		for cursor := 0; cursor < len(args); cursor++ {
			bboltconnection.RemoveLocation(args[cursor])
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

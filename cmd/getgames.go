/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"d3u/bboltconnection"
	"github.com/spf13/cobra"
)

// getgamesCmd represents the getgames command
var getgamesCmd = &cobra.Command{
	Use:   "getgames",
	Short: "List games",
	Long:  `List all games with their location info`,
	Run: func(cmd *cobra.Command, args []string) {
		bboltconnection.GetGames()
	},
}

func init() {
	rootCmd.AddCommand(getgamesCmd)
}

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bboltconnection.GetGames()
	},
}

func init() {
	rootCmd.AddCommand(getgamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getgamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getgamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

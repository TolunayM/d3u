package cmd

import (
	"d3u/bboltconnection"
	"github.com/spf13/cobra"
)

// rmCmd represents the remove command
var rmCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove game data",
	Long: `Remove game's data from database.
You can type game's name right after remove command 
Example:

			d3u remove Cyberpunk2077

but i suggest you to using double quotes because some games has spaces in their name.

For example when you use remove command with game like "The Last of Us", commandline thinks every word is a different game.
You can just use double quotes like

			d3u remove "The Last of Us"

Actually Tlou has an exe like tlou-i.exe but you got the concept.
And you can use for removing multiple games' at once.

			d3u remove Cyberpunk2077 "The Last of Us" "Diablo IV"
`,
	Run: func(cmd *cobra.Command, args []string) {

		for cursor := 0; cursor < len(args); cursor++ {
			bboltconnection.RemoveLocation(args[cursor])
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

package cmd

import (
	"d3u/bboltconnection"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqweek/dialog"
	"path/filepath"
)

// addgameCmd represents the addgame command
var addgameCmd = &cobra.Command{
	Use:   "addgame",
	Short: "Adding games directories",
	Long:  `This for adding your games directories for updating dlss`,
	Run: func(cmd *cobra.Command, args []string) {

		file, err := dialog.File().Filter(".exe", "exe").Title("Select a file").Load()
		if err != nil {
			fmt.Println("Error selecting file:", err)
			return
		}

		absolutePath, err := filepath.Abs(file)
		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return
		}

		gameDirectory := filepath.Dir(absolutePath)
		gameName := filepath.Base(absolutePath)
		bboltconnection.Dbinit(gameName, gameDirectory)
		fmt.Println("Game directory added to bboltconnection", gameDirectory, gameName)
	},
}

func init() {
	rootCmd.AddCommand(addgameCmd)
}

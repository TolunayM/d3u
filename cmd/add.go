package cmd

import (
	"d3u/bboltconnection"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqweek/dialog"
	"path/filepath"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding games' directories to database",
	Long: `You need to select game's or application's executable (.exe) file to save directory to database. 
This doesn't alter executables just saves executables directory for upgrading or downgrading to DLSS files.`,
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
		gameName = gameName[0 : len(gameName)-4]
		bboltconnection.Addgame(gameName, gameDirectory)
		fmt.Println("Game directory added to database", gameDirectory, gameName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

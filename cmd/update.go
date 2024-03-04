package cmd

import (
	"d3u/tools"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// updateCmd represents the update command

var home, _ = os.UserHomeDir()
var d3uHomeDB = home + "\\d3u\\db"
var d3uHomeDLSS = home + "\\d3u\\dlss"
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Upgrade or downgrade your dlss versions.",
	Long: `Upgrade or downgrade your dlss versions.
You can type game's name right after update command 
Example:

			d3u update Cyberpunk2077

but i suggest you to using double quotes because some games has spaces in their name.

For example when you use update command with game like "The Last of Us", commandline thinks every word is a different game.
You can just use double quotes like

			d3u update "The Last of Us"

Actually Tlou has an exe like tlou-i.exe but you got the concept.
And you can use for updating multiple games' at once.

			d3u update Cyberpunk2077 "The Last of Us" "Diablo IV"

After that you can specify your DLSS version with -version or -v flag default is latest DLSS version.
Example:

			d3u update Cyberpunk2077 -v "3.5.0" or d3u update Cyberpunk2077 -v 3.5.0

just be sure you typed correct version of a DLSS. After that it will automatically checks current DLSS version and it'll decide to upgrade or downgrade.
`,
	Run: func(cmd *cobra.Command, args []string) {

		version, _ := cmd.Flags().GetString("version")
		customDLL := d3uHomeDLSS + "\\nvngx_dlss_" + version + ".dll"

		/*
			this flag was used for get the games that wanted to be updated
			but after thinking about that I decide to change.
			on cobra-cli all commands actually flags of root command
			and that means you can already have parse flag as named "args".
			this way is better for "command line philosophy"
		*/
		//gameSelection := cmd.Flags().Parse("game")

		db, err := bolt.Open(d3uHomeDB+"\\my.bboltconnection", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// check before download if requested dlss file is downloaded before
		file, err := os.ReadFile(customDLL)
		if err != nil {
			fmt.Println("Dlss files is not present locally download starting...")
			tools.DownloadDLSS(version)
			file, _ = os.ReadFile(customDLL)
		}
		// updating
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			c := b.Cursor()

			if len(args) != 0 {
				for cursor := 0; cursor < len(args); cursor++ {
					if string(b.Get([]byte(args[cursor]))) != "" {
						if tools.CheckDlssVersion(string(b.Get([]byte(args[cursor])))) != version {
							err = os.WriteFile(string(b.Get([]byte(args[cursor])))+"\\nvngx_dlss.dll", file, 0644)
							fmt.Println(args[cursor] + " Updated to " + version + " Successfully")
						} else {
							fmt.Println("Same version no need to update")
						}

					} else {
						fmt.Println("Game is not present in db check your typo or use get command for listing games added to database")
					}

				}

			} else {

				for key, value := c.First(); key != nil; key, value = c.Next() {

					if tools.CheckDlssVersion(string(b.Get(key))) != version {
						fmt.Println(string(key) + " Updated to " + version + " Successfully")
						err = os.WriteFile(string(value)+"\\nvngx_dlss.dll", file, 0644)
					} else {
						fmt.Println(string(key) + " already has " + version + " no need to update")
					}

					if err != nil {
						fmt.Println(err)
					}

				}

			}
			return nil
		})

	},
}

func init() {

	updateCmd.Flags().StringP("version", "v", "latest", "Version specifier")
	rootCmd.AddCommand(updateCmd)
}

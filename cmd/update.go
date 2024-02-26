/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// updateCmd represents the update command

var home, _ = os.UserHomeDir()
var d3uHomeDB = home + "\\d3u\\db"
var d3uHomeDLSS = home + "\\d3u\\dlss"
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `Update dlss`,
	Run: func(cmd *cobra.Command, args []string) {

		version, _ := cmd.Flags().GetString("version")
		customDLL := d3uHomeDLSS + "\\nvngx_dlss_" + version + ".dll"

		/*
			this flag was used for get the games that wanted to be updated
			but after thinking about that I decide to change
			on cobra-cli all commands actually flags of root command
			and that means you can already have parse flag as named "args"
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
			downloadDLSS(version)
			file, _ = os.ReadFile(customDLL)
		}

		// updating
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			c := b.Cursor()

			if len(args) != 0 {
				for cursor := 0; cursor < len(args); cursor++ {
					if string(b.Get([]byte(args[cursor]))) != "" {
						err = os.WriteFile(string(b.Get([]byte(args[cursor])))+"\\nvngx_dlss.dll", file, 0644)
						fmt.Println(args[cursor] + " Updated to " + version + " Successfully")
					} else {
						fmt.Println("Game is not present check your typo or use getgames command for listing games added to database")
					}

				}

			} else {

				for key, value := c.First(); key != nil; key, value = c.Next() {

					fmt.Println(string(key) + " Updated to " + version + " Successfully")
					err = os.WriteFile(string(value)+"\\nvngx_dlss.dll", file, 0644)

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
	//updateCmd.Flags().StringP("game", "g", "", "Game selection for updating specific games")

	rootCmd.AddCommand(updateCmd)
}

func downloadDLSS(version string) {

	downloadLink := "https://github.com/TolunayM/dlss-repo/releases/download/" + version + "/nvngx_dlss.dll"
	customDLL := d3uHomeDLSS + "\\nvngx_dlss_" + version + ".dll"

	fmt.Println("Downloading dlss files.\nThis may take a minute based on your connection speed.")
	download := exec.Command(
		"curl",
		"-o",
		customDLL,
		"-L",
		downloadLink)
	_, err := download.Output()

	if err != nil {
		_ = fmt.Errorf("something happened %s", err)
	}

}

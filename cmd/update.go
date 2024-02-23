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
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `Update dlss`,
	Run: func(cmd *cobra.Command, args []string) {

		version, _ := cmd.Flags().GetString("version")
		gameSelection, _ := cmd.Flags().GetString("game")

		downloadLink := "https://github.com/TolunayM/dlss-repo/releases/download/" + version + "/nvngx_dlss.dll"
		db, err := bolt.Open("my.bboltconnection", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		fmt.Println("Downloading dlss files.\nThis may take a minute based on your connection speed.")

		download := exec.Command(
			"curl",
			"-OL",
			downloadLink)
		_, err = download.Output()

		if err != nil {
			_ = fmt.Errorf("Something happened %s", err)
		}

		file, err := os.ReadFile("nvngx_dlss.dll")
		if err != nil {
			fmt.Println(err)
		}

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			c := b.Cursor()

			for key, value := c.First(); key != nil; key, value = c.Next() {
				if gameSelection != "" {

					fmt.Println(gameSelection + " Updated to " + version + " Successfully")
					err = os.WriteFile(string(b.Get([]byte(gameSelection)))+"\\nvngx_dlss.dll", file, 0644)

					return nil

				} else {

					fmt.Println(string(key) + " Updated to " + version + " Successfully")
					err = os.WriteFile(string(value)+"\\nvngx_dlss.dll", file, 0644)

				}

				if err != nil {
					fmt.Println(err)
				}

			}
			return nil
		})

	},
}

func init() {

	updateCmd.Flags().StringP("version", "v", "latest", "Version specifier")
	updateCmd.Flags().StringP("game", "g", "", "Game selection for updating specific games")
	rootCmd.AddCommand(updateCmd)
}

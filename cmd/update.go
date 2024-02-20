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

		db, err := bolt.Open("my.bboltconnection", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		fmt.Println("Downloading dlss files.\nThis may take a minute based on your connection speed.")

		download := exec.Command(
			"curl",
			"-OL",
			"https://github.com/TolunayM/dlss-repo/releases/download/3.5.0/nvngx_dlss.dll")
		_, err = download.Output()

		if err != nil {
			_ = fmt.Errorf("Something happened %s", err)
		}

		//request, _ := http.NewRequest("GET", "https://github.com/TolunayM/dlss-repo/releases/download/3.5.10/nvngx_dlss.dll", nil)
		//response, _ := http.DefaultClient.Do(request)
		//
		//defer func(Body io.ReadCloser) {
		//	err := Body.Close()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}(response.Body)
		//
		//f, _ := os.OpenFile("nvngx_dlss_3.5.10.dll", os.O_CREATE|os.O_WRONLY, 0644)
		//defer f.Close()
		//
		//bar := progressbar.DefaultBytes(
		//	response.ContentLength,
		//	"Downloading",
		//)
		//
		//io.Copy(io.MultiWriter(f, bar), response.Body)

		file, err := os.ReadFile("nvngx_dlss.dll")

		if err != nil {
			fmt.Println(err)
		}

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("MyBucket"))
			c := b.Cursor()
			for key, value := c.First(); key != nil; key, value = c.Next() {

				fmt.Println(string(key) + " Updated Successfully")
				err = os.WriteFile(string(value)+"\\nvngx_dlss.dll", file, 0644)
				if err != nil {
					fmt.Println(err)
				}
				//err := os.Rename("path", string(value)+"\\nvngx_dlss.dll")
				//
				//if err != nil {
				//	fmt.Println(err)
				//}
			}

			return nil
		})

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

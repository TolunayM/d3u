/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `Update dlss`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Downloading dlss files.\nThis may take a minute based on your connection speed.")
		//
		//download := exec.Command("curl", "-OL", "===download link here===")
		//_, err := download.Output()
		//
		//if err != nil {
		//	_ = fmt.Errorf("Something happened %s", err)
		//}

		request, _ := http.NewRequest("GET", "===downloadlinkhere===", nil)
		response, _ := http.DefaultClient.Do(request)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(response.Body)

		f, _ := os.OpenFile("nvngx_dlss_3.5.10.zip", os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		bar := progressbar.DefaultBytes(
			response.ContentLength,
			"downloading",
		)

		io.Copy(io.MultiWriter(f, bar), response.Body)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

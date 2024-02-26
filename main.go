package main

import (
	"d3u/cmd"
	"fmt"
	"os"
)

func main() {

	var home, _ = os.UserHomeDir()
	d3uHomeDB := home + "\\d3u\\db"
	d3uHomeDLSS := home + "\\d3u\\dlss"
	if _, err := os.Stat(d3uHomeDB); os.IsNotExist(err) {
		err := os.Mkdir(d3uHomeDB, os.ModeDir)
		if err != nil {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(d3uHomeDLSS); os.IsNotExist(err) {
		err := os.Mkdir(d3uHomeDLSS, os.ModeDir)
		if err != nil {
			fmt.Println(err)
		}
	}
	cmd.Execute()

}

package main

import (
	"fmt"
	"github.com/sqweek/dialog"
	"os"
	"path/filepath"
)

func changeDir() {
	file, err := dialog.File().Filter(".exe", "exe").Title("Select a file").Load()
	if err != nil {
		fmt.Println("Error selecting file:", err)
		return
	}

	// Get the absolute path of the selected file
	absolutePath, err := filepath.Abs(file)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	workingDirectory := filepath.Dir(absolutePath)
	err = os.Chdir(workingDirectory)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Print the new current working directory
	fmt.Println("Changed directory to:", workingDirectory)
}

func main() {

}

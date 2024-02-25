/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var rmCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove game location",
	Long:  `Remove game's location from db'`,
	Run: func(cmd *cobra.Command, args []string) {

		for cursor := 0; cursor < len(args); cursor++ {
			removeLocation(args[cursor])
			fmt.Println(args[cursor] + " Removed Succesfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func removeLocation(removedFile string) {

	//open db connection
	db, err := bolt.Open(home+"\\my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Delete([]byte(removedFile))
		return err
	})
}

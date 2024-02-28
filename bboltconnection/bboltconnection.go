package bboltconnection

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

var home, _ = os.UserHomeDir()
var d3uHomeDB = home + "\\d3u\\db"

func Addgame(gameName string, gameDirectory string) {

	db, err := bolt.Open(d3uHomeDB+"\\my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		return nil
	})
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte(gameName), []byte(gameDirectory))
		if err != nil {
			return fmt.Errorf("Transaction: %s", err)
		}
		return nil
	})

}

func GetGames() {

	db, err := bolt.Open(d3uHomeDB+"\\my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()

		for key, value := c.First(); key != nil; key, value = c.Next() {

			fmt.Printf("Game = %s, location = %s\n", string(key), string(value))
		}
		return nil
	})
}

func RemoveLocation(removedData string) {

	//open db connection
	db, err := bolt.Open(d3uHomeDB+"\\my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		if b.Get([]byte(removedData)) != nil {
			_ = b.Delete([]byte(removedData))
			fmt.Println(removedData + " Removed Succesfully")
		} else {
			fmt.Println("Game is not present in db check your typo or use get command for listing games added to database")
		}

		return err
	})
	if err != nil {
		return
	}
}

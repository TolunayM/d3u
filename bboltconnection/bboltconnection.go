package bboltconnection

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// change init to main and dont create a new db everytime
func Dbinit(gameName string, gameDirectory string) {

	db, err := bolt.Open("my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	defer db.Close()

	//TODO create distinct add func
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte(gameName), []byte(gameDirectory))
		if err != nil {
			return fmt.Errorf("transaction: %s", err)
		}
		return nil
	})

}

// TODO fix
func GetGames() {

	db, err := bolt.Open("my.bboltconnection", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for key, value := c.First(); key != nil; key, value = c.Next() {

			//fmt.Println(c)
			fmt.Printf("Game = %s, location = %s\n", string(key), string(value))
		}

		//v := b.Get([]byte(gameName))
		//fmt.Printf("Location of a :%s is %s\n", gameName, v)
		return nil
	})
}

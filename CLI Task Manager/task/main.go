package main

import (
	"clitaskmanager/task/cmd"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
	defer db.Close()
}

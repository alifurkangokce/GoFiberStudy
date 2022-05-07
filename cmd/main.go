package main

import (
	"fmt"
	"log"
	"study/db"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	fmt.Print(database.Config)
}

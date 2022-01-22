package main

import (
	"fmt"
	"./db"
)

func main() {
	fmt.Println("Connecting to AWS MongoDB")
	client := db.dbConnection()
}

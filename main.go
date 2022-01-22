package main

import (
	"fmt"

	db "reamer-chess-db/db"
)

func main() {
	fmt.Println("Connecting to AWS MongoDB")
	client := db.FormConnection()

}

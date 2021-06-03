package main

import (
	"handleVM/client"
	"handleVM/handledb"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := handledb.Connect()
	client.RunClient(db)
	defer db.Close()
}

package main

import (
	"context"
	"handleVM/client"
	"handleVM/handledb"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	db, err := handledb.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	client.RunClient(ctx, db)
}

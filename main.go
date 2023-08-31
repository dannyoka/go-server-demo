package main

import (
	"fmt"

	"github.com/dannyoka/go-server/internal/db"
	"github.com/dannyoka/go-server/internal/server"
)

func main() {
	fmt.Println("Hello world")
	client, err := db.InitDB(); if err != nil {
		panic(err)
	}

	
	server.InitRouter(client)
}

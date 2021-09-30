package main

import (
	"pg-conn/src/server"

	"pg-conn/src/database"
)

func main() {
	database.Migrate()
	server := server.NewServer()

	server.Run()

}

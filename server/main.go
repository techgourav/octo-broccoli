package main

import (
	"octo-broccoli/server/src/db"
	"octo-broccoli/server/src/server"
)

func main() {
	db.Init()
	server.Start()
}

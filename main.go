package main

import (
	"autochess/db"
	"autochess/server"
)

func main() {
		db.Init()
		server.Init()
}

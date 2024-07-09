package main

import (
	"github.com/Eirc-lab-star/apiServer/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server.Start()
}

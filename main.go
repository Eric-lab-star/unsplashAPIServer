package main

import (
	"fmt"

	"github.com/Eirc-lab-star/apiServer/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("go is runing")
	server.Start()
}

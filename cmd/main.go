package main

import (
	"Todo/cmd/server"
)

func main() {

	go server.StartServer()

	select {}

}

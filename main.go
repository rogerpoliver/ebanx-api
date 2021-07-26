package main

import "github.com/rogerpoliver/ebanx-api/server"

func main() {
	server := server.CreateServer()
	server.StartServer()
}

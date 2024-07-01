package main

import (
	"sponsor-sv/api"
	"log"
)

func main() {
	log.Println("Hello, this is VARMETA")
	log.Println("Starting server")
	defaultHost := "localhost:8765"
	api.StartH2CServer(defaultHost)
}

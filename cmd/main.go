package main

import (
	"sponsor-sv/router"
	"log"
)

func main() {
	log.Println("Hello, this is varmeta")
	log.Println("Starting server")
	defaultHost := "localhost:8765"
	router.StartH2CServer(defaultHost)
}

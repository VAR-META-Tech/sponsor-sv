package main

import (
	"log"
	"sponsor-sv/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello, this is VARMETA")
	log.Println("Starting server")
	defaultHost := "localhost:8765"
	api.StartH2CServer(defaultHost)
}

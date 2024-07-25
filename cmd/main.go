package main

import (
	"log"
	"sponsor-sv/api"
	"sponsor-sv/configs"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello, this is VARMETA")
	log.Println("Starting server")
	env, err := configs.GetEnv("./configs/config.yaml")
	if err != nil {
		panic(err)
	}
	api.StartH2CServer(env.Host + ":" + env.Port)
}

package api

import (

	// "go-microservices-training/services/handler"

	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	log.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}
func StartH2CServer(host string) {
	h2s := &http2.Server{}

	handler := NewHandler()
	server := &http.Server{
		Addr:    host,
		Handler: h2c.NewHandler(handler, h2s),
	}

	log.Printf("Listening [%s]...\n", host)
	go checkErr(server.ListenAndServe(), "while listening")
}

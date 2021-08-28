package main

import (
	"golang_tdd/src/app/server"
	"log"
	"net/http"
)

func main() {
	server := server.NewPlayerServer(server.Dependencies{
		Store: server.NewInMemoryStore(),
	})
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":5000", server))
}

package main

import (
	"log"
	"net/http"

	"golang_tdd/src/app/server"
)

func main() {
	server := server.PlayerServer{}
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":5000", &server))
}

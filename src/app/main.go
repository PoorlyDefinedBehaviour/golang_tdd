package main

import (
	"golang_tdd/src/app/server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":5000", handler))
}

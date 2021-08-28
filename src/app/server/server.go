package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (server *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	playerID := strings.TrimPrefix(r.URL.Path, "players/")
	playerID = strings.TrimSuffix(playerID, "/score")

	fmt.Fprint(w, server.store.GetPlayerScore(playerID))
}

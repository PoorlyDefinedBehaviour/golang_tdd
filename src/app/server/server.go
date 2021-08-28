package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

func (server *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		server.processPlayerWin(w, r)
	case http.MethodGet:
		server.getPlayerScore(w, r)
	}
}

func (server *PlayerServer) getPlayerScore(w http.ResponseWriter, r *http.Request) {
	playerID := strings.TrimPrefix(r.URL.Path, "players/")
	playerID = strings.TrimSuffix(playerID, "/score")

	score, found := server.store.GetPlayerScore(playerID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (server *PlayerServer) processPlayerWin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	return
}

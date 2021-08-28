package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
	RecordWin(id string)
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

func getPlayerIDFromPath(path string) string {
	playerID := strings.TrimPrefix(path, "players/")
	playerID = strings.TrimSuffix(playerID, "/score")
	return playerID
}

func (server *PlayerServer) getPlayerScore(w http.ResponseWriter, r *http.Request) {
	playerID := getPlayerIDFromPath(r.URL.Path)

	score, found := server.store.GetPlayerScore(playerID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, score)
}

func (server *PlayerServer) processPlayerWin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	playerID := getPlayerIDFromPath(r.URL.Path)
	server.store.RecordWin(playerID)
}

package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type InMemoryStore struct {
	store map[string]int
	mutex sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		store: make(map[string]int),
		mutex: sync.RWMutex{},
	}
}

func (store *InMemoryStore) GetPlayerScore(id string) (int, bool) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	score, found := store.store[id]
	return score, found
}

func (store *InMemoryStore) RecordWin(id string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	score, ok := store.store[id]
	if !ok {
		store.store[id] = 1
	} else {
		store.store[id] = score + 1
	}
}

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
	RecordWin(id string)
}

type PlayerServer struct {
	store PlayerStore
}

type Dependencies struct {
	Store PlayerStore
}

func NewPlayerServer(deps Dependencies) *PlayerServer {
	return &PlayerServer{
		store: deps.Store,
	}
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

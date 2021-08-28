package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PlayerStoreStub struct {
	scores   map[string]int
	winCalls []string
}

func (store *PlayerStoreStub) GetPlayerScore(id string) (int, bool) {
	score, found := store.scores[id]
	return score, found
}

func (store *PlayerStoreStub) RecordWin(id string) {
	store.winCalls = append(store.winCalls, id)
}

func TestGetPlayerScore(t *testing.T) {
	t.Parallel()

	server := PlayerServer{
		store: &PlayerStoreStub{
			scores: map[string]int{
				"1": 10,
				"2": 20,
			},
		},
	}

	t.Run("returns player 1 score", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "players/1/score", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "10", response.Body.String())
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("returns player 2 score", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "players/2/score", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "20", response.Body.String())
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "players/9999999/score", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}

func TestStorePlayerWins(t *testing.T) {
	t.Parallel()

	store := PlayerStoreStub{
		scores:   make(map[string]int),
		winCalls: make([]string, 0),
	}

	server := PlayerServer{
		store: &store,
	}

	t.Run("records player wins", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodPost, "players/1/score", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, []string{"1"}, store.winCalls)

		assert.Equal(t, http.StatusAccepted, response.Code)
	})
}

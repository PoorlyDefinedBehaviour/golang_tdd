package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PlayerStoreStub struct {
	scores map[string]int
}

func (store *PlayerStoreStub) GetPlayerScore(id string) int {
	return store.scores[id]
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
	})

	t.Run("returns player 2 score", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "players/2/score", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "20", response.Body.String())
	})
}

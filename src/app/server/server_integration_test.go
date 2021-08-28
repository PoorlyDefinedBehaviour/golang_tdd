package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	t.Parallel()

	store := NewInMemoryStore()

	server := PlayerServer{store: store}

	request, _ := http.NewRequest(http.MethodPost, "players/1/score", nil)

	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)

	response := httptest.NewRecorder()

	request, _ = http.NewRequest(http.MethodGet, "players/1/score", nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, "3", response.Body.String())
}

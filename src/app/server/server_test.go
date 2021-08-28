package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayerScore(t *testing.T) {
	t.Parallel()

	request, _ := http.NewRequest(http.MethodGet, "players/1/score", nil)

	response := httptest.NewRecorder()

	PlayerServer(response, request)

	assert.Equal(t, "20", response.Body.String())
}

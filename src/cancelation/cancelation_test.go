package cancelation

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	response string
	canceled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		data := "hello, world"

		store := StubStore{response: data}

		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, data, response.Body.String())
		assert.False(t, store.canceled)
	})

	t.Run("tells store to cancel work if request is canceled", func(t *testing.T) {
		t.Parallel()

		data := "hello, world"

		store := StubStore{response: data}

		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.True(t, store.canceled)
	})
}

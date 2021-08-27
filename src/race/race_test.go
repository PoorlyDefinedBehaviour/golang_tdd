package race

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRace(t *testing.T) {
	t.Parallel()

	t.Run("requests urls in parallel and returns the first url to respond with non error response", func(t *testing.T) {
		t.Parallel()

		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
		defer slowServer.Close()

		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer fastServer.Close()

		expected := fastServer.URL

		actual, err := First([]string{slowServer.URL, fastServer.URL})

		assert.Nil(t, err)

		assert.Equal(t, expected, actual)
	})

	t.Run("returns error if no url returns before the timeout", func(t *testing.T) {
		t.Parallel()

		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Second)
			w.WriteHeader(http.StatusOK)
		}))

		_, err := FirstWithTimeout([]string{slowServer.URL}, 1*time.Millisecond)

		assert.True(t, errors.Is(err, ErrTimedOut))
	})
}

package race

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// makes concurrenct GET requests to the urls and
// returns the the first that returns a non error response
func First(urls []string) (string, error) {
	return FirstWithTimeout(urls, 10*time.Second)
}

var ErrTimedOut = errors.New("timed out")

func FirstWithTimeout(urls []string, timeout time.Duration) (string, error) {
	first := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for _, url := range urls {
		url := url

		go func() {
			request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				return
			}
			defer response.Body.Close()

			if response.StatusCode < 200 || response.StatusCode >= 500 {
				return
			}

			first <- url
		}()
	}

	select {
	case firstURLToRespond := <-first:
		return firstURLToRespond, nil
	case <-time.After(timeout):
		return "", errors.WithStack(ErrTimedOut)
	}
}

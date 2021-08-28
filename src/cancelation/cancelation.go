package cancelation

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case out := <-data:
			fmt.Fprint(w, out)
		case <-r.Context().Done():
			store.Cancel()
		}
	}
}

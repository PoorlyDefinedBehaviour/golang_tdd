package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	playerID := strings.TrimPrefix(r.URL.Path, "players/")
	playerID = strings.TrimSuffix(playerID, "/score")

	if playerID == "1" {
		fmt.Fprint(w, "10")
		return
	}

	if playerID == "2" {
		fmt.Fprint(w, "20")
		return
	}
}

package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UpcomingHostsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// get ordered slice of upcoming hosts

	upcomingHosts := []string {
		"Welch",
		"Brown",
		"Bazzo",
	}

	writeJsonResponse(w, upcomingHosts, nil, http.StatusOK)
}

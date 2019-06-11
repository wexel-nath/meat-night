package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func UpcomingHostsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	upcomingHosts := logic.GetUpcomingHosts()
	writeJsonResponse(w, upcomingHosts, nil, http.StatusOK)
}

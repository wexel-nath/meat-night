package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func ListMateosHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	method := ps.ByName("method")
	mateos, err := logic.GetAllMateos(method)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, mateos, messages, http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, mateos, nil, http.StatusOK)
}

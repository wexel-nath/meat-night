package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func ListMateosHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	mateos, err := logic.GetAllMateos()
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, mateos, messages, http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, mateos, nil, http.StatusOK)
}

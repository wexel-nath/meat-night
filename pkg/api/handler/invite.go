package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func AcceptHostInviteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inviteID := ps.ByName("inviteID")

	mateo, err := logic.GetMateoByInviteID(inviteID)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusUnprocessableEntity)
		return
	}

	// get venue somehow!

	// create dinner

	//newDinner, err := logic.CreateDinner(dinner)
	//if err != nil {
	//	logger.Error(err)
	//	messages := []string { err.Error() }
	//	writeJsonResponse(w, nil, messages, http.StatusUnprocessableEntity)
	//	return
	//}

	logger.Info("mateo[%s] accepted the invite[%s] to host", mateo.LastName, inviteID)

	writeJsonResponse(w, nil, nil, http.StatusCreated)
}

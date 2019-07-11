package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func AcceptHostInviteHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	inviteID := ps.ByName("inviteID")
	err := logic.AcceptHostInvite(inviteID)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, nil, nil, http.StatusOK)
}

func DeclineHostInviteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inviteID := ps.ByName("inviteID")
	err := logic.DeclineHostInvite(inviteID)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	writeJsonResponse(w, nil, nil, http.StatusOK)
}

func AcceptGuestInviteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inviteID := ps.ByName("inviteID")

	mateo, err := logic.GetMateoByInviteID(inviteID)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusUnprocessableEntity)
		return
	}

	// update invite record to 'accepted'

	// insert a guest entry

	logger.Info("mateo[%s] accepted the invite[%s] to attend", mateo.LastName, inviteID)

	writeJsonResponse(w, nil, nil, http.StatusCreated)
}

func DeclineGuestInviteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inviteID := ps.ByName("inviteID")

	mateo, err := logic.GetMateoByInviteID(inviteID)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusUnprocessableEntity)
		return
	}

	// update invite record to 'declined'

	logger.Info("mateo[%s] declined the invite[%s] to attend", mateo.LastName, inviteID)

	writeJsonResponse(w, nil, nil, http.StatusCreated)
}

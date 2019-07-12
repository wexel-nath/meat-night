package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

// TODO: these handlers should return some html or redirect to a page

func AcceptHostInvite(r *http.Request, ps httprouter.Params) (interface{}, int, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.AcceptHostInvite(inviteID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return nil, http.StatusOK, nil
}

func DeclineHostInvite(r *http.Request, ps httprouter.Params) (interface{}, int, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.DeclineHostInvite(inviteID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return nil, http.StatusOK, err
}

func AcceptGuestInvite(_ *http.Request, ps httprouter.Params) (interface{}, int, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.AcceptGuestInvite(inviteID)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return nil, http.StatusOK, err
}

func DeclineGuestInvite(_ *http.Request, ps httprouter.Params) (interface{}, int, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.DeclineGuestInvite(inviteID)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return nil, http.StatusOK, err
}

package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func AcceptHostInvite(ps httprouter.Params) (string, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.AcceptHostInvite(inviteID)
	return "thumbs-up", err
}

func DeclineHostInvite(ps httprouter.Params) (string, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.DeclineHostInvite(inviteID)
	return "thumbs-down", err
}

func AcceptGuestInvite(ps httprouter.Params) (string, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.RespondToGuestInvite(inviteID, model.TypeInviteAccepted)
	return "thumbs-up", err
}

func DeclineGuestInvite(ps httprouter.Params) (string, error) {
	inviteID := ps.ByName("inviteID")
	err := logic.RespondToGuestInvite(inviteID, model.TypeInviteDeclined)
	return "thumbs-down", err
}

package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func inviteHost(mateoID int64) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, nil)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

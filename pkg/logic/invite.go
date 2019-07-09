package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func inviteHost(mateoID int64) (model.Invite, error) {
	// create unique ID
	uniqueID := "TEST-ID-1"

	row, err := database.InsertInvite(uniqueID, model.TypeInviteHost, mateoID, nil)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

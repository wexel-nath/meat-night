package logic

import (
	"time"

	"github.com/pkg/errors"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func inviteHost(mateoID int64) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	nextDinnerTime := getNextDinnerTime(time.Now())

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, nil, &nextDinnerTime)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func getNextDinnerTime(now time.Time) time.Time {
	dayDiff := config.GetDinnerDay() - now.Weekday()
	if dayDiff < 0 || now.Hour() > 19 {
		dayDiff += 7
	}

	d := now.AddDate(0, 0, int(dayDiff))
	return time.Date(d.Year(), d.Month(), d.Day(), 19, 0, 0, 0, time.Local)
}

func inviteGuest(mateoID int64, dinnerID int64) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, &dinnerID, nil)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func getInviteByID(inviteID string) (model.Invite, error) {
	row, err := database.SelectInviteByID(inviteID)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func acceptInvite(inviteID string) (model.Invite, error) {
	row, err := database.UpdateInvite(inviteID, model.TypeInviteAccepted)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func declineInvite(inviteID string) (model.Invite, error) {
	row, err := database.UpdateInvite(inviteID, model.TypeInviteDeclined)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func AcceptHostInvite(inviteID string) error {
	logger.Info("host invite[%s] has been accepted", inviteID)

	invite, err := validateInvite(getInviteByID(inviteID))
	if err != nil {
		return err
	}

	mateo, err := GetMateoByInviteID(inviteID)
	if err != nil {
		return err
	}

	// todo: get venue somehow!
	dinner, err := CreateDinner(
		invite.DinnerTime.Format(model.DateFormat),
		"PLACEHOLDER",
		mateo.LastName,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = acceptInvite(inviteID)
	if err != nil {
		return err
	}

	return alertGuestsForDinner(mateo, dinner.ID)
}

func DeclineHostInvite(inviteID string) error {
	logger.Info("host invite[%s] has been accepted", inviteID)

	_, err := validateInvite(getInviteByID(inviteID))
	if err != nil {
		return err
	}

	mateo, err := GetMateoByInviteID(inviteID)
	if err != nil {
		return err
	}

	_, err = declineInvite(inviteID)
	if err != nil {
		return err
	}

	_, err = inviteNextHost(mateo)
	return err
}

func inviteNextHost(host model.Mateo) (model.Invite, error) {
	mateos, err := GetAllMateos(model.TypeLegacy)
	if err != nil {
		return model.Invite{}, err
	}

	for index, mateo := range mateos {
		if mateo.ID == host.ID && index < len(mateos) - 1 {
			nextHost := mateos[index + 1]
			return inviteHost(nextHost.ID)
		}
	}

	return model.Invite{}, errors.New("cannot find a host to invite")
}

func validateInvite(invite model.Invite, err error) (model.Invite, error) {
	if err != nil {
		return invite, err
	}
	if invite.InviteStatus != model.TypeInvitePending {
		return invite, model.ErrInviteHasResponse
	}
	if time.Until(invite.DinnerTime) < 6 * time.Hour {
		return invite, model.ErrInviteLateResponse
	}
	return invite, nil
}

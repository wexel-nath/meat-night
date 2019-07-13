package logic

import (
	"time"

	"github.com/pkg/errors"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func createHostInvite(mateoID int64) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	nextDinnerTime := getNextDinnerTime(time.Now())

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, nil, nextDinnerTime)
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

func createGuestInvite(mateoID int64, dinnerID int64, dinnerTime time.Time) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	row, err := database.InsertInvite(inviteID, model.TypeInviteGuest, mateoID, &dinnerID, dinnerTime)
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
		[]string{ mateo.LastName },
	)
	if err != nil {
		return err
	}

	_, err = acceptInvite(inviteID)
	if err != nil {
		return err
	}

	return inviteGuests(mateo, dinner.ID, invite.DinnerTime)
}

func DeclineHostInvite(inviteID string) error {
	logger.Info("host invite[%s] has been declined", inviteID)

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

	nextHost, err := findNextHost(mateo)
	if err != nil {
		return err
	}

	return inviteMateoToHost(nextHost)
}

func findNextHost(host model.Mateo) (model.Mateo, error) {
	mateos, err := GetAllMateos(model.TypeLegacy)
	if err != nil {
		return model.Mateo{}, err
	}

	for index, mateo := range mateos {
		if mateo.ID == host.ID && index < len(mateos) - 1 {
			return mateos[index + 1], nil
		}
	}

	return model.Mateo{}, errors.New("cannot find a host to invite")
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

func RespondToGuestInvite(inviteID string, response string) error {
	logger.Info("guest invite[%s] has been %s", inviteID, response)

	invite, err := validateInvite(getInviteByID(inviteID))
	if err != nil {
		return err
	}

	if response == model.TypeInviteAccepted {
		mateo, err := GetMateoByInviteID(inviteID)
		if err != nil {
			return err
		}

		err = database.InsertGuest(invite.DinnerID, mateo.ID)
		if err != nil {
			return err
		}
	}

	_, err = database.UpdateInvite(inviteID, response)
	if err != nil {
		return err
	}

	invites, err := getInvitesForDinner(invite.DinnerID)
	if err != nil {
		return err
	}
	if !hasPendingInvites(invites) {
		return SendGuestListEmail(true)
	}

	return nil
}

func hasPendingInvites(invites []model.Invite) bool {
	for _, invite := range invites {
		if invite.InviteStatus == model.TypeInvitePending {
			return true
		}
	}
	return false
}

func getInvitesForDinner(dinnerID int64) ([]model.Invite, error) {
	rows, err := database.SelectAllInvitesForDinner(dinnerID)
	if err != nil {
		return nil, err
	}

	invites := make([]model.Invite, 0)
	for _, row := range rows {
		invite, err := model.NewInviteFromRow(row)
		if err != nil {
			return invites, err
		}
		invites = append(invites, invite)
	}

	return invites, nil
}

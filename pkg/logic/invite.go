package logic

import (
	"fmt"
	"time"

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

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, nil)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func getUpcomingDinnerDate() (time.Time, error) {
	now := time.Now()

	today := now.Weekday()

	dinnerDay := config.GetDinnerDay()

	if today >= dinnerDay {
		return now, fmt.Errorf("too late to invite a host. today[%s] dinnerDay[%s]", today.String(), dinnerDay.String())
	}

	dinnerDate := now.AddDate(0, 0, int(dinnerDay - today))
	return dinnerDate, nil
}

func acceptInvite(inviteID string) (model.Invite, error) {
	row, err := database.UpdateInvite(inviteID, model.TypeInviteAccepted)
	if err != nil {
		return model.Invite{}, err
	}

	return model.NewInviteFromRow(row)
}

func AcceptHostInvite(inviteID string) error {

	// todo: check if invite is pending & valid?
	mateo, err := GetMateoByInviteID(inviteID)
	if err != nil {
		return err
	}

	logger.Info("mateo[%s] accepted the invite[%s] to host", mateo.LastName, inviteID)

	// todo: get venue somehow!

	// todo: get date of upcoming meat night

	// create dinner
	d := model.Dinner{
		Date:  "01-01-01",
		Venue: "PLACEHOLDER",
		Host:  mateo.LastName,
	}

	_, err = CreateDinner(d)
	if err != nil {
		return err
	}

	// update invite record to 'accepted'
	_, err = acceptInvite(inviteID)
	if err != nil {
		return err
	}

	// invite other mateos to dinner
	// todo: alert guests needs a dinner id for the invite link
	return alertGuestsForDinner(mateo)
}

package logic

import (
	"errors"
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

func AcceptHostInvite(inviteID string) error {
	invite, err := getInviteByID(inviteID)
	if err != nil {
		return err
	}

	if time.Until(invite.DinnerTime) < 6 * time.Hour {
		return errors.New("it's too late to accept an invite to host")
	}

	mateo, err := GetMateoByInviteID(inviteID)
	if err != nil {
		return err
	}

	logger.Info("mateo[%s] accepted the invite[%s] to host", mateo.LastName, inviteID)

	// todo: get venue somehow!

	d := model.Dinner{
		Date:  invite.DinnerTime.Format(model.DateFormat),
		Venue: "PLACEHOLDER",
		Host:  mateo.LastName,
	}

	dinner, err := CreateDinner(d)
	if err != nil {
		return err
	}

	_, err = acceptInvite(inviteID)
	if err != nil {
		return err
	}

	return alertGuestsForDinner(mateo, dinner.ID)
}

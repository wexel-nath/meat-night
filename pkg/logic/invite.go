package logic

import (
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

func inviteGuest(mateoID int64, dinnerID int64) (model.Invite, error) {
	inviteID, err := generateUniqueID(mateoID)
	if err != nil {
		return model.Invite{}, err
	}

	row, err := database.InsertInvite(inviteID, model.TypeInviteHost, mateoID, &dinnerID)
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

	// todo: check if invite is pending & valid?
	// accept needs to happen at least 8? hours before meat night

	mateo, err := GetMateoByInviteID(inviteID)
	if err != nil {
		return err
	}

	logger.Info("mateo[%s] accepted the invite[%s] to host", mateo.LastName, inviteID)

	// todo: get venue somehow!

	dinnerDate, err := getUpcomingDinnerDate()
	if err != nil {
		return err
	}

	d := model.Dinner{
		Date:  dinnerDate.Format(model.DateFormat),
		Venue: "PLACEHOLDER",
		Host:  mateo.LastName,
	}

	dinner, err := CreateDinner(d)
	if err != nil {
		return err
	}

	// update invite record to 'accepted'
	_, err = acceptInvite(inviteID)
	if err != nil {
		return err
	}

	return alertGuestsForDinner(mateo, dinner.ID)
}

func getUpcomingDinnerDate() (time.Time, error) {
	now := time.Now()
	today := now.Weekday()
	dinnerDay := config.GetDinnerDay()

	//if today > dinnerDay {
	//	return now, fmt.Errorf("too late to invite a host. today[%s] dinnerDay[%s]", today.String(), dinnerDay.String())
	//}

	dinnerDate := now.AddDate(0, 0, int(dinnerDay - today))
	return dinnerDate, nil
}

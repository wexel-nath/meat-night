package logic

import (
	"time"

	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func SendInviteHostEmail() error {
	mateos, err := GetAllMateos(model.TypeLegacy)
	if err != nil {
		return err
	}

	return inviteMateoToHost(mateos[0])
}

func inviteMateoToHost(mateo model.Mateo) error {
	invite, err := createHostInvite(mateo.ID)
	if err != nil {
		return err
	}

	emailFunc := func() error {
		return email.SendAlertHostEmail(mateo, invite.InviteID)
	}
	return maybeSendEmail(mateo.ID, model.TypeInviteHost, emailFunc, false)
}

func inviteGuests(hostMateo model.Mateo, dinnerID int64, dinnerTime time.Time) error {
	logger.Info("Sending guest emails for mateo[%s] dinner", hostMateo.LastName)

	guests, err := GetAllMateosExceptHost(hostMateo.ID)
	if err != nil {
		return err
	}

	for _, guest := range guests {
		err = sendInviteGuestEmail(guest, hostMateo.FirstName, dinnerID, dinnerTime)
		if err != nil {
			logger.Error(err)
		}
	}

	return nil
}

func sendInviteGuestEmail(guest model.Mateo, hostName string, dinnerID int64, dinnerTime time.Time) error {
	// dont email Jimbo
	if guest.FirstName == "James"{
		return nil
	}

	invite, err := createGuestInvite(guest.ID, dinnerID, dinnerTime)
	if err != nil {
		return err
	}

	emailFunc := func() error {
		return email.SendAlertGuestEmail(guest, hostName, invite.InviteID)
	}
	return maybeSendEmail(guest.ID, model.TypeInviteGuest, emailFunc, false)
}

func SendGuestListEmail(forceSend bool) error {
	dinner, err := GetLatestDinner()
	if err != nil {
		return err
	}

	mateo, err := GetMateoByLastName(dinner.Host)
	if err != nil {
		return err
	}

	invites, err := getInvitesForDinner(dinner.ID)
	if err != nil {
		return err
	}

	invitees := map[string][]string{
		"accepted": { "You" },
		"declined": {},
		"pending":  {},
	}
	for _, invite :=  range invites {
		invitedMateo, err := GetMateoByID(invite.MateoID)
		if err != nil {
			return err
		}
		invitees[invite.InviteStatus] = append(invitees[invite.InviteStatus], invitedMateo.FirstName)
	}

	emailFunc := func() error {
		return email.SendGuestListEmail(mateo, dinner.ID, invitees)
	}

	return maybeSendEmail(mateo.ID, model.TypeGuestList, emailFunc, forceSend)
}

func maybeSendEmail(mateoID int64, messageType string, emailFunc func() error, force bool) error {
	logger.Info("Sending mateo[%d] a message[%s]", mateoID, messageType)

	if !force {
		//check if mateo has received email recently
		messages, err := getRecentMessagesForMateo(mateoID, messageType)
		if err != nil {
			return err
		}
		if len(messages) > 0 {
			logger.Info("mateo[%d] has received a %s message recently, not sending", mateoID, messageType)
			return nil
		}
	}

	err := emailFunc()
	if err != nil {
		return err
	}

	_, err = database.InsertMessage(messageType, mateoID)
	return err
}

func getRecentMessagesForMateo(mateoID int64, messageType string) ([]model.Message, error) {
	messages := make([]model.Message, 0)

	rows, err := database.SelectRecentMessagesForMateoEvent(messageType, mateoID)
	if err != nil {
		return messages, err
	}

	for _, row := range rows {
		m, err := model.NewMessageFromMap(row)
		if err != nil {
			return messages, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

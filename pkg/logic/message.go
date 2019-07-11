package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func AlertHost() error {
	mateos, err := GetAllMateos(model.TypeLegacy)
	if err != nil {
		return err
	}
	mateoToAlert := mateos[0]

	invite, err := inviteHost(mateoToAlert.ID)
	if err != nil {
		return err
	}

	emailFunc := func() error {
		return email.SendAlertHostEmail(mateoToAlert, invite.InviteID)
	}
	return maybeSendEmail(mateoToAlert.ID, model.TypeAlertHost, emailFunc)
}

func alertGuestsForDinner(hostMateo model.Mateo, dinnerID int64) error {
	logger.Info("Sending guest emails for mateo[%s] dinner", hostMateo.LastName)

	guests, err := GetAllMateosExceptHost(hostMateo.ID)
	if err != nil {
		return err
	}

	for _, guest := range guests {
		err = sendAlertGuestEmail(guest, hostMateo.FirstName, dinnerID)
		if err != nil {
			logger.Error(err)
		}
	}

	return nil
}

func sendAlertGuestEmail(guest model.Mateo, hostName string, dinnerID int64) error {
	// dont email Jimbo
	if guest.FirstName == "James"{
		return nil
	}

	invite, err := inviteGuest(guest.ID, dinnerID)
	if err != nil {
		return err
	}

	emailFunc := func() error {
		return email.SendAlertGuestEmail(guest, hostName, invite.InviteID)
	}
	return maybeSendEmail(guest.ID, model.TypeAlertGuest, emailFunc)
}

func maybeSendEmail(mateoID int64, messageType string, emailFunc func() error) error {
	logger.Info("Sending mateo[%d] a message[%s]", mateoID, messageType)

	// check if mateo has received email recently
	//messages, err := getRecentMessagesForMateo(mateoID, messageType)
	//if err != nil {
	//	return err
	//}
	//if len(messages) > 0 {
	//	logger.Info("mateo[%d] has received a %s message recently, not sending", mateoID, messageType)
	//	return nil
	//}

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

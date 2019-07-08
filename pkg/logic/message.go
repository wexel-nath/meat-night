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

	emailFunc := func() error {
		return email.SendAlertHostEmail(mateoToAlert)
	}
	return maybeSendEmail(mateoToAlert.ID, model.TypeAlertHost, emailFunc)
}

func AlertGuests() error {
	mateos, err := GetAllMateos(model.TypeLegacy)
	if err != nil {
		return err
	}
	hostMateo := mateos[0]

	mateos, err = GetAllMateos("")
	if err != nil {
		return err
	}

	for _, mateo := range mateos {
		// dont email the host or Jimbo
		if mateo.ID == hostMateo.ID || mateo.FirstName == "James"{
			continue
		}

		emailFunc := func() error {
			return email.SendAlertGuestEmail(mateo, hostMateo.FirstName)
		}
		err = maybeSendEmail(mateo.ID, model.TypeAlertGuest, emailFunc)
		if err != nil {
			return err
		}
	}

	return nil
}

func maybeSendEmail(mateoID int64, messageType string, emailFunc func() error) error {
	logger.Info("Sending mateo[%d] a %s message", mateoID, messageType)

	// check if mateo has received email recently
	messages, err := getRecentMessagesForMateo(mateoID, messageType)
	if err != nil {
		return err
	}
	if len(messages) > 0 {
		logger.Info("mateo[%d] has received a %s message recently, not sending", mateoID, messageType)
		return nil
	}

	err = emailFunc()
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

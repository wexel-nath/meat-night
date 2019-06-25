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
	logger.Info("Sending %s a %s message", mateoToAlert.LastName, model.TypeAlertHost)

	// check if mateo has received email recently
	messages, err := GetRecentAlertHostMessagesForMateo(mateoToAlert.ID)
	if err != nil {
		return err
	}
	if len(messages) > 0 {
		logger.Info(
			"%s has received a %s message recently, not sending",
			mateoToAlert.LastName,
			model.TypeAlertHost,
		)
		return nil
	}

	err = email.SendAlertHostEmail(mateoToAlert)
	if err != nil {
		return err
	}

	_, err = database.InsertMessage(model.TypeAlertHost, mateoToAlert.ID)
	return err
}

func GetRecentAlertHostMessagesForMateo(mateoID int64) ([]model.Message, error) {
	messages := make([]model.Message, 0)

	rows, err := database.SelectRecentMessagesForMateoEvent(model.TypeAlertHost, mateoID)
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

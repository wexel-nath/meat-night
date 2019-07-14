package logic

import (
	"time"

	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

// TODO: put these in a transaction so we can rollback if guests do not get inserted
func CreateDinner(date string, venue string, hostName string, attendeeNames []string) (model.Dinner, error) {
	dateTime, err := time.Parse(model.DateFormat, date)
	if err != nil {
		return model.Dinner{}, err
	}

	row, err := database.InsertDinner(dateTime, venue, hostName)
	if err != nil {
		return model.Dinner{}, err
	}

	dinner, err := model.NewDinnerFromMap(row)
	if err != nil {
		return dinner, err
	}

	dinner.Attended, err = database.InsertGuests(dinner.ID, attendeeNames)
	return dinner, err
}

func GetAllDinners() ([]model.Dinner, error) {
	rows, err := database.SelectAllDinners()
	if err != nil {
		return nil, err
	}

	dinners := make([]model.Dinner, 0)
	for _, row := range rows {
		dinner, err := model.NewDinnerFromMap(row)
		if err != nil {
			return dinners, err
		}

		dinners = append(dinners, dinner)
	}

	return dinners, nil
}

func GetLatestDinner() (model.Dinner, error) {
	row, err := database.SelectLatestDinner()
	if err != nil {
		return model.Dinner{}, err
	}

	return model.NewDinnerFromMap(row)
}

func UpdateDinnerVenue(dinnerID int64, venue string) error {
	logger.Info("Updating dinner[%d] with venue[%s]", dinnerID, venue)
	return database.UpdateDinnerVenue(dinnerID, venue)
}

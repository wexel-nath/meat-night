package logic

import (
	"time"

	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

// TODO: put these in a transaction so we can rollback if guests do not get inserted
func CreateDinner(dinner model.Dinner) (model.Dinner, error) {
	date, err := time.Parse(model.DateFormat, dinner.Date)
	if err != nil {
		return dinner, err
	}

	row, err := database.InsertDinner(date, dinner.Venue, dinner.Host)
	if err != nil {
		return dinner, err
	}

	newDinner, err := model.NewDinnerFromMap(row)
	if err != nil {
		return newDinner, err
	}

	newDinner.Attended, err = database.InsertGuests(newDinner.ID, dinner.Attended)
	return newDinner, err
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

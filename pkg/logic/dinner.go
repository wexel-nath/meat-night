package logic

import (
	"time"

	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

// TODO: Validate dinner request
func CreateDinner(dinner model.DinnerRequestDto) error {
	date, err := time.Parse(model.DateFormat, dinner.Date)
	if err != nil {
		return err
	}

	// TODO: put these in a transaction so we can rollback if guests do not get inserted
	dinnerID, err := database.InsertDinner(date, dinner.HostID, dinner.Venue)
	if err != nil {
		return err
	}

	return database.InsertGuests(dinnerID, dinner.GuestIDs)
}
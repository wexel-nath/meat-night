package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func GetAllMateos(method string) ([]model.Mateo, error) {
	var rows []map[string]interface{}
	var err error

	if method == model.TypeLegacy {
		rows, err = database.SelectAllMateosLegacy()
	} else {
		rows, err = database.SelectAllMateos()
	}

	return getMateosFromRows(rows, err)
}

func getMateosFromRows(rows []map[string]interface{}, err error) ([]model.Mateo,error) {
	if err != nil {
		return nil, err
	}

	mateos := make([]model.Mateo, 0)
	for _, row := range rows {
		mateo, err := model.NewMateoFromMap(row)
		if err != nil {
			return mateos, err
		}

		mateos = append(mateos, mateo)
	}

	return mateos, nil
}

func GetMateoByInviteID(inviteID string) (model.Mateo, error) {
	row, err := database.SelectMateoByInviteID(inviteID)
	if err != nil {
		return model.Mateo{}, err
	}

	return model.NewMateoFromMap(row)
}

func GetAllMateosExceptHost(hostID int64) ([]model.Mateo, error) {
	return getMateosFromRows(database.SelectAllMateosExceptHost(hostID))
}

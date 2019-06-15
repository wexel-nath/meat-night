package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func GetAllMateos() ([]model.Mateo, error) {
	return getMateosFromRows(database.SelectAllMateos())
}

func GetAllMateosLegacy() ([]model.Mateo, error) {
	return getMateosFromRows(database.SelectAllMateosLegacy())
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

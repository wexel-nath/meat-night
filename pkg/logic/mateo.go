package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func GetAllMateos() ([]model.Mateo, error) {
	rows, err := database.SelectAllMateos()
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

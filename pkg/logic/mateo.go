package logic

import (
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func GetAllMateos(method string) ([]model.Mateo, error) {
	var rows []map[string]interface{}
	var err error

	if method == model.TypeLegacy {
		rows, err = database.SelectAllMateosLegacy()

		message := email.Create(
			"test@example.com",
			"Test Mailgun Email",
			"This is a test email",
			"nathanwelch_@hotmail.com",
		)
		err := email.Send(message)
		logger.LogIfErr(err)
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

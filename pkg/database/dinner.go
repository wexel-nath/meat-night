package database

import (
	"strings"
	"time"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func InsertDinner(date time.Time, mateoID int64, venue string) (int64, error) {
	query := `
		INSERT INTO dinner(
			date, mateo_id, venue
		)
		VALUES
			($1, $2, $3)
		RETURNING 
			id
	`
	params := []interface{}{
		date,
		mateoID,
		venue,
	}

	db := getConnection()
	var id int64
	err := db.QueryRow(query, params...).Scan(&id)
	return id, err
}

func SelectAllDinners() ([]map[string]interface{}, error) {
	columns := model.GetDinnerColumns()
	query := `
		SELECT ` + strings.Join(columns, ", ") + `
		FROM     dinner
			JOIN mateo USING (mateo_id)
		ORDER BY date
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, columns)
}

package database

import (
	"strings"
	"time"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func SelectAllDinners() ([]map[string]interface{}, error) {
	columns := model.GetDinnerColumns()
	query := `
		SELECT
			` + strings.Join(columns, ", ") + `
		FROM
			dinner
			JOIN mateo USING (mateo_id)
		ORDER BY
			date
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, columns)
}

func SelectLatestDinner() (map[string]interface{}, error) {
	columns := model.GetDinnerColumns()
	query := `
		SELECT
			` + strings.Join(columns, ", ") + `
		FROM
			dinner
			JOIN mateo USING (mateo_id)
		ORDER BY
			date DESC,
			dinner_id DESC
		LIMIT
			1
	`

	db := getConnection()
	row := db.QueryRow(query)
	return scanRowToMap(row, columns)
}

func InsertDinner(date time.Time, venue string, lastName string) (map[string]interface{}, error) {
	columns := model.GetDinnerColumns()
	query := `
		WITH insert_dinner AS (
			INSERT INTO dinner (date, venue, mateo_id)
				SELECT $1, $2, mateo_id
				FROM mateo
				WHERE last_name = $3
				LIMIT 1
			RETURNING *
		)
		SELECT
			` + strings.Join(columns, ", ") + `
		FROM
			insert_dinner
			JOIN mateo USING (mateo_id)
	`

	db := getConnection()
	row := db.QueryRow(query, date, venue, lastName)
	return scanRowToMap(row, columns)
}

func UpdateDinnerVenue(dinnerID int64, venue string) error {
	query := `
		UPDATE
			dinner
		SET
			venue = $2
		WHERE
			dinner_id = $1
			AND venue = 'PLACEHOLDER'
	`

	db := getConnection()
	_, err := db.Exec(query, dinnerID, venue)
	return err
}

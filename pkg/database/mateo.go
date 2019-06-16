package database

import (
	"strings"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func SelectAllMateos() ([]map[string]interface{}, error) {
	columns := model.GetMateoColumns()
	query := `
		WITH guest_counts AS (
			SELECT   mateo_id, COUNT(*) AS total_attended
			FROM     guest
			GROUP BY mateo_id
		),
		host_counts AS (
			SELECT   mateo_id, COUNT(*) AS total_hosted
			FROM     dinner
			GROUP BY mateo_id
		)
		SELECT 
			` + strings.Join(columns, ", ") + `
		FROM
			mateo
			LEFT JOIN guest_counts USING (mateo_id)
			LEFT JOIN host_counts USING (mateo_id)
		ORDER BY
			mateo_id
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, columns)
}

func SelectAllMateosLegacy() ([]map[string]interface{}, error) {
	columns := model.GetMateoLegacyColumns()
	query := `
		WITH last_host AS (
			SELECT   mateo_id, MAX(date) AS last_host_date
			FROM     dinner
			GROUP BY mateo_id
		)
		SELECT 
			` + strings.Join(columns, ", ") + `
		FROM
			mateo
			JOIN guest USING (mateo_id)
			JOIN dinner USING (dinner_id)
			JOIN last_host USING (mateo_id)
		WHERE
			dinner.date > last_host_date
		GROUP BY
			mateo_id, last_host_date
		ORDER BY
			attended DESC, last_host_date
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, columns)
}

func SelectMateoByLastName(lastName string) (map[string]interface{}, error) {
	return nil, nil
}

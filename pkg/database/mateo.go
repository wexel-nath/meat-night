package database

import (
	"strings"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func SelectAllMateos() ([]map[string]interface{}, error) {
	columns := model.GetMateoSortColumns()
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
	query := `
		WITH last_host AS (
			SELECT   mateo_id, MAX(date) AS last_host_date
			FROM     dinner
			GROUP BY mateo_id
		)
		SELECT
			mateo.mateo_id AS mateo_id,
			first_name,
			last_name,
			last_host_date,
			COUNT(*) AS attended
		FROM
			mateo
			JOIN guest USING (mateo_id)
			JOIN dinner USING (dinner_id)
			JOIN last_host ON last_host.mateo_id = guest.mateo_id
		WHERE
			dinner.date > last_host_date
		GROUP BY
			mateo.mateo_id, last_host_date
		ORDER BY
			attended DESC, last_host_date;
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, model.GetMateoSortLegacyColumns())
}

func SelectMateoByLastName(lastName string) (map[string]interface{}, error) {
	columns := model.GetMateoColumns()
	query := `
		SELECT ` + strings.Join(columns, ", ") + `
		FROM   mateo
		WHERE  last_name = $1
	`

	db := getConnection()
	row := db.QueryRow(query, lastName)
	return scanRowToMap(row, columns)
}

package database

import (
	"strings"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func SelectAllMateos() ([]map[string]interface{}, error) {
	columns := model.GetMateoColumns()
	query := `
		WITH guest_counts AS (
			SELECT   guest_id, COUNT(*) AS guest_count
			FROM     guest
			GROUP BY guest_id
		),
		host_counts AS (
			SELECT   host_id, COUNT(*) AS host_count
			FROM     dinner
			GROUP BY host_id
		)
		SELECT ` + strings.Join(columns, ", ") + `
		FROM      mateo m
			LEFT JOIN guest_counts gc ON m.id = gc.guest_id
			LEFT JOIN host_counts hc ON m.id = hc.host_id
		ORDER BY  m.id
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
			SELECT   host_id, MAX(date) AS last_host_date
			FROM     dinner
			GROUP BY host_id
		)
		SELECT ` + strings.Join(columns, ", ") + `
		FROM   mateo
			JOIN guest ON guest.guest_id = mateo.id
			JOIN dinner ON dinner.id = guest.dinner_id
			JOIN last_host ON last_host.host_id = mateo.id
		WHERE  dinner.date > last_host_date
		GROUP BY mateo.id, last_host_date
		ORDER BY attended DESC, last_host_date
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, columns)
}

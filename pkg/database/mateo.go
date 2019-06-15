package database

import "strings"

var (
	mateoColumns = []string {
		"id",
		"first_name",
		"last_name",
		"guest_count",
		"host_count",
	}

	mateoLegacyColumns = []string {
		"id",
		"first_name",
		"last_name",
		"last_host_date",
		"COUNT(*) AS attended",
	}
)

func SelectAllMateos() ([]map[string]interface{}, error) {
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
		SELECT ` + strings.Join(mateoColumns, ", ") + `
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

	return scanRowsToMap(rows, mateoColumns)
}

func SelectAllMateosLegacy() ([]map[string]interface{}, error) {
	query := `
		WITH host_turn AS (
			SELECT   host_id, MAX(date) AS last_host_date
			FROM     dinner
			GROUP BY host_id
		)
		SELECT ` + strings.Join(mateoLegacyColumns, ", ") + `
		FROM   mateo
			JOIN guest ON guest.guest_id = mateo.id
			JOIN dinner ON dinner.id = guest.dinner_id
			JOIN host_turn ON host_turn.host_id = mateo.id
		WHERE  dinner.date > last_host_date
		GROUP BY mateo.id, last_host_date
		ORDER BY attended DESC, last_host_date
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, mateoLegacyColumns)
}

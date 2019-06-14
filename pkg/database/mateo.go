package database

import "strings"

var (
	mateoColumns = []string {
		"id",
		"first_name",
		"last_name",
		"guest_count",
		"host_count",
		"guest_count::NUMERIC / host_count AS guest_ratio",
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
		ORDER BY  m.id;
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, mateoColumns)
}

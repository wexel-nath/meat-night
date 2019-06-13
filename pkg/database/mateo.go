package database

import "strings"

var (
	mateoColumns = []string {
		"id",
		"first_name",
		"last_name",
	}
)

func SelectAllMateos() ([]map[string]interface{}, error) {
	query := `
		SELECT 
			` + strings.Join(mateoColumns, ", ") + `
		FROM
			mateo
		ORDER BY
			last_name
	`

	db := getConnection()
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return scanRowsToMap(rows, mateoColumns)
}

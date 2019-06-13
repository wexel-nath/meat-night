package database

import "strings"

func SelectUpcomingHosts() []string {
	db := getConnection()

	query := `
		SELECT
			` + strings.Join(mateoColumns, ", ") + `
		FROM 
			mateo
			JOIN dinner ON mateo.id = dinner.host_id
	`

	db.Query(query)

	return []string {
		"Welch",
		"Brown",
		"Bazzo",
	}
}

package database

import "strings"

func SelectUpcomingHosts() []string {
	db := getConnection()

	query := `
		SELECT
			` + strings.Join(mateoColumns, ", ") + `
		FROM 
			mateo
			JOIN meat_night ON mateo.id = meat_night.host_id
	`

	db.Query(query)

	return []string {
		"Welch",
		"Brown",
		"Bazzo",
	}
}

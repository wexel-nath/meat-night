package database

func SelectUpcomingHosts() []string {
	db := getConnection()

	query := `
		SELECT first_name + ' ' + last_name
		FROM   mateo
	`

	db.Query(query)

	return []string {
		"Welch",
		"Brown",
		"Bazzo",
	}
}

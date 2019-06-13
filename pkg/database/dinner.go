package database

import "time"

func InsertDinner(date time.Time, hostID int64, venue string) (int64, error) {
	query := `
		INSERT INTO dinner(
			date, host_id, venue
		)
		VALUES
			($1, $2, $3)
		RETURNING 
			id
	`
	params := []interface{}{
		date,
		hostID,
		venue,
	}

	db := getConnection()
	var id int64
	err := db.QueryRow(query, params...).Scan(&id)
	return id, err
}

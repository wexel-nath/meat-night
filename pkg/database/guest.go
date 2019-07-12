package database

import (
	"github.com/lib/pq"
)

func InsertGuests(dinnerID int64, lastNames []string) ([]string, error) {
	query := `
		WITH insert_guests AS (
			INSERT INTO guest (dinner_id, mateo_id)
				SELECT $1, mateo_id
				FROM   mateo
				WHERE  last_name = ANY($2)
			RETURNING *
		)
		SELECT last_name
		FROM
			insert_guests
			JOIN mateo USING (mateo_id)
	`

	db := getConnection()
	rows, err := db.Query(query, dinnerID, pq.Array(lastNames))
	if err != nil {
		return nil, err
	}

	return scanSingleColumnToStringSlice(rows)
}

func InsertGuest(dinnerID int64, mateoID int64) error {
	query := `
		INSERT INTO guest (dinner_id, mateo_id)
		VALUES ($1, $2)
	`

	db := getConnection()
	_, err := db.Exec(query, dinnerID, mateoID)
	return err
}

func SelectAllGuestsForDinner(dinnerID int64) ([]string, error) {
	query := `
		SELECT
			first_name
		FROM
			mateo
			JOIN guest USING (mateo_id)
		WHERE
			dinner_id = $1
	`

	db := getConnection()
	rows, err := db.Query(query, dinnerID)
	if err != nil {
		return nil, err
	}

	return scanSingleColumnToStringSlice(rows)
}

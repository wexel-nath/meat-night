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

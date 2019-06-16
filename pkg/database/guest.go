package database

import (
	"fmt"

	"github.com/lib/pq"
)

func InsertGuests(dinnerID int64, mateoIDs []int64) error {
	query := buildInsertGuestsQuery(len(mateoIDs))

	ids := make([]interface{}, len(mateoIDs))
	for i, id := range mateoIDs {
		ids[i] = id
	}
	params := append([]interface{}{ dinnerID }, ids...)

	db := getConnection()
	_, err := db.Exec(query, params...)
	return err
}

func buildInsertGuestsQuery(numGuests int) string {
	query := `
		INSERT INTO guest(
			dinner_id, mateo_id
		)
		VALUES
	`

	for i := 1; i <= numGuests; i++ {
		format := "($1, $%d)"
		if i != numGuests {
			format += ","
		}
		query += fmt.Sprintf(format, i + 1)
	}

	return query
}

func InsertGuestsByLastName(dinnerID int64, lastNames []string) ([]string, error) {
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

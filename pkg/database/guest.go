package database

import "fmt"

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

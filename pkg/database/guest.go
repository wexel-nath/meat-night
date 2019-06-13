package database

import "fmt"

func InsertGuests(dinnerID int64, guestIDs []int64) error {
	query := buildInsertGuestsQuery(len(guestIDs))

	ids := make([]interface{}, len(guestIDs))
	for i, v := range guestIDs {
		ids[i] = v
	}
	params := append([]interface{}{ dinnerID }, ids...)

	db := getConnection()
	_, err := db.Exec(query, params...)
	return err
}

func buildInsertGuestsQuery(numGuests int) string {
	query := `
		INSERT INTO guest(
			dinner_id, guest_id
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

package database

import (
	"strings"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func InsertMessage(event string, mateoID int64) (map[string]interface{}, error) {
	columns := model.GetMessageColumns()
	query := `
		INSERT INTO message (event, mateo_id)
		VALUES ($1, $2)
		RETURNING ` + strings.Join(columns, ", ")

	db := getConnection()
	row := db.QueryRow(query, event, mateoID)
	return scanRowToMap(row, columns)
}

func SelectRecentMessagesForMateoEvent(event string, mateoID int64) ([]map[string]interface{}, error) {
	columns := model.GetMessageColumns()
	query := `
		SELECT ` + strings.Join(columns, ", ") + `
		FROM
			message
		WHERE
			message_event = $1
			AND mateo_id = $2
			AND message_timestamp > NOW() - INTERVAL '4 days'
	`

	db := getConnection()
	rows, err := db.Query(query, event, mateoID)
	if err != nil {
		return nil, err
	}
	return scanRowsToMap(rows, columns)
}

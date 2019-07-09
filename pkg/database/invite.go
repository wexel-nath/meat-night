package database

import (
	"strings"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func InsertInvite(inviteID string, inviteType string, mateoID int64, dinnerID *int64) (map[string]interface{}, error) {
	columns := model.GetInviteColumns()
	query := `
		INSERT INTO invite (invite_id, invite_type, mateo_id, dinner_id)
		VALUES ($1, $2, $3, $4)
		RETURNING ` + strings.Join(columns, ", ")

	db := getConnection()
	row := db.QueryRow(query, inviteID, inviteType, mateoID, dinnerID)
	return scanRowToMap(row, columns)
}

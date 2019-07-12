package database

import (
	"strings"
	"time"

	"github.com/wexel-nath/meat-night/pkg/model"
)

func InsertInvite(
	inviteID string,
	inviteType string,
	mateoID int64,
	dinnerID *int64,
	dinnerTime time.Time,
) (map[string]interface{}, error) {
	columns := model.GetInviteColumns()
	query := `
		INSERT INTO invite (
			invite_id,
			invite_type,
			mateo_id,
			dinner_id,
			dinner_time
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING
			` + strings.Join(columns, ", ")

	db := getConnection()
	row := db.QueryRow(query, inviteID, inviteType, mateoID, dinnerID, dinnerTime)
	return scanRowToMap(row, columns)
}

func SelectInviteByID(inviteID string) (map[string]interface{}, error) {
	columns := model.GetInviteColumns()
	query := `
		SELECT
			` + strings.Join(columns, ", ") + `
		FROM
			invite
		WHERE
			invite_id = $1
	`

	db := getConnection()
	row := db.QueryRow(query, inviteID)
	return scanRowToMap(row, columns)
}

func SelectMateoByInviteID(inviteID string) (map[string]interface{}, error) {
	columns := model.GetMateoColumns()
	query := `
		SELECT
			` +  strings.Join(columns, ", ") + `
		FROM
			mateo
			JOIN invite USING (mateo_id)
		WHERE
			invite_id = $1
	`

	db := getConnection()
	row := db.QueryRow(query, inviteID)
	return scanRowToMap(row, columns)
}

func UpdateInvite(inviteID string, inviteStatus string) (map[string]interface{}, error) {
	columns := model.GetInviteColumns()
	query := `
		UPDATE
			invite
		SET
			invite_status = $1
		WHERE
			invite_id = $2
		RETURNING
			` + strings.Join(columns, ", ")

	db := getConnection()
	row := db.QueryRow(query, inviteStatus, inviteID)
	return scanRowToMap(row, columns)
}

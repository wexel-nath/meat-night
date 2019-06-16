package database

import (
	"database/sql"
	"database/sql/driver"
	"github.com/wexel-nath/meat-night/pkg/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	SelectAllMateosExpectedQuery = `
		WITH guest_counts AS \(
			SELECT   guest_id, COUNT\(\*\) AS guest_count
			FROM     guest
			GROUP BY guest_id
		\),
		host_counts AS \(
			SELECT   host_id, COUNT\(\*\) AS host_count
			FROM     dinner
			GROUP BY host_id
		\)
		SELECT
			id,
			first_name,
			last_name,
			guest_count,
			host_count
		FROM      mateo m
			LEFT JOIN guest_counts gc ON m.id = gc.guest_id
			LEFT JOIN host_counts hc ON m.id = hc.host_id
		ORDER BY  m.id
	`

	SelectAllMateosLegacyExpectedQuery = `
		WITH last_host AS \(
			SELECT   host_id, MAX\(date\) AS last_host_date
			FROM     dinner
			GROUP BY host_id
		\)
		SELECT
			mateo.id AS id,
			first_name,
			last_name,
			last_host_date,
			COUNT\(\*\) AS attended
		FROM   mateo
			JOIN guest ON guest.guest_id = mateo.id
			JOIN dinner ON dinner.id = guest.dinner_id
			JOIN last_host ON last_host.host_id = mateo.id
		WHERE  dinner.date > last_host_date
		GROUP BY mateo.id, last_host_date
		ORDER BY attended DESC, last_host_date
	`
)

type MockRow []driver.Value

type Mock struct {
	ExpectQuery string
	ExpectRows  []MockRow
	ExpectErr   error
}

func GetMockDB(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	setMockConnection(db)
	return mock
}

// setMockConnection is used for testing only
func setMockConnection(db *sql.DB) {
	connection = db
}


func GetValues(mateo model.Mateo, sort string) []driver.Value {
	values := []driver.Value{
		mateo.ID,
		mateo.FirstName,
		mateo.LastName,
	}

	if sort == model.TypeLegacy {
		lastHostDate, _ := time.Parse(model.DateFormat, mateo.LastHostDate)
		values = append(
			values,
			lastHostDate,
			mateo.Attended,
		)
	} else {
		values = append(
			values,
			mateo.GuestCount,
			mateo.HostCount,
		)
	}

	return values
}

func GetMap(mateo model.Mateo, sort string) map[string]interface{} {
	m := map[string]interface{}{
		"id":          mateo.ID,
		"first_name":  mateo.FirstName,
		"last_name":   mateo.LastName,
	}

	if sort == model.TypeLegacy {
		m["last_host_date"] = mateo.LastHostDate
		m["attended"] = mateo.Attended
	} else {
		m["guest_count"] = mateo.GuestCount
		m["host_count"] = mateo.HostCount
	}

	return m
}

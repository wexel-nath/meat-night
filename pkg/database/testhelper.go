package database

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	SelectAllMateosExpectedQuery = `
		WITH guest_counts AS \(
			SELECT   mateo_id, COUNT\(\*\) AS total_attended
			FROM     guest
			GROUP BY mateo_id
		\),
		host_counts AS \(
			SELECT   mateo_id, COUNT\(\*\) AS total_hosted
			FROM     dinner
			GROUP BY mateo_id
		\)
		SELECT
			mateo_id,
			first_name,
			last_name,
			total_attended,
			total_hosted
		FROM
			mateo
			LEFT JOIN guest_counts USING \(mateo_id\)
			LEFT JOIN host_counts USING \(mateo_id\)
		ORDER BY
			mateo_id
	`

	SelectAllMateosLegacyExpectedQuery = `
		WITH last_host AS \(
			SELECT   mateo_id, MAX\(date\) AS last_host_date
			FROM     dinner
			GROUP BY mateo_id
		\)
		SELECT
			mateo.mateo_id AS mateo_id,
			first_name,
			last_name,
			last_host_date,
			COUNT\(\*\) AS attended
		FROM
			mateo
			JOIN guest USING \(mateo.id\)
			JOIN dinner USING \(dinner_id\)
			JOIN last_host ON last_host.mateo_id = guest.mateo_id
		WHERE
			dinner.date > last_host_date
		GROUP BY
			mateo.mateo_id, last_host_date
		ORDER BY
			attended DESC, last_host_date
	`
)

var (
	TestJohn = model.Mateo{
		ID:            1,
		FirstName:     "John",
		LastName:      "Doe",
		TotalAttended: 5,
		TotalHosted:   2,
		GuestRatio:    2.5,
	}
	TestAdam = model.Mateo{
		ID:            2,
		FirstName:     "Adam",
		LastName:      "Samuel",
		TotalAttended: 4,
		TotalHosted:   2,
		GuestRatio:    2.0,
	}

	// Legacy
	TestBob = model.Mateo{
		ID:           3,
		FirstName:    "Bob",
		LastName:     "Jane",
		LastHostDate: "12-05-19",
		Attended:     4,
	}
	TestDavid = model.Mateo{
		ID: 4,
		FirstName:    "David",
		LastName:     "Wilson",
		LastHostDate: "01-06-19",
		Attended:     1,
	}
)

type MockRow []driver.Value

type Mock struct {
	ExpectQuery   string
	ExpectColumns []string
	ExpectRows    []MockRow
	ExpectErr     error
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
			mateo.TotalAttended,
			mateo.TotalHosted,
		)
	}

	return values
}

func GetMap(mateo model.Mateo, sort string) map[string]interface{} {
	m := map[string]interface{}{
		"mateo_id":   mateo.ID,
		"first_name": mateo.FirstName,
		"last_name":  mateo.LastName,
	}

	if sort == model.TypeLegacy {
		m["last_host_date"] = mateo.LastHostDate
		m["attended"] = mateo.Attended
	} else {
		m["total_attended"] = mateo.TotalAttended
		m["total_hosted"] = mateo.TotalHosted
	}

	return m
}

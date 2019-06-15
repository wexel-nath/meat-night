package database

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wexel-nath/meat-night/pkg/model"
)

type mockRow []driver.Value

type mock struct {
	expectRows []mockRow
	expectErr  error
}

func getMockDB(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	SetMockConnection(db)
	return mock
}

const (
	selectAllMateosExpectedQuery = `
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
)

func TestSelectAllMateos(t *testing.T) {
	type want struct {
		mateoMap []map[string]interface{}
		err      error
	}

	tests := []struct {
		name string
		mock
		want
	}{
		{
			name: "Single Row Returned",
			mock: mock{
				expectRows: []mockRow{
					model.TestJohn.ToRow(),
				},
			},
			want: want{
				mateoMap: []map[string]interface{}{
					model.TestJohn.ToMap(),
				},
				err: nil,
			},
		},
		{
			name: "Multiple Rows Returned",
			mock: mock{
				expectRows: []mockRow{
					model.TestJohn.ToRow(),
					model.TestAdam.ToRow(),
				},
			},
			want: want{
				mateoMap: []map[string]interface{}{
					model.TestJohn.ToMap(),
					model.TestAdam.ToMap(),
				},
				err: nil,
			},
		},
		{
			name: "Connection Error",
			mock: mock{
				expectErr: errors.New("connection error"),
			},
			want: want{
				mateoMap: nil,
				err:      errors.New("connection error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			dbMock := getMockDB(st)
			query := dbMock.ExpectQuery(selectAllMateosExpectedQuery)

			if test.mock.expectErr != nil {
				query.WillReturnError(test.mock.expectErr)
			} else {
				mockRows := sqlmock.NewRows(mateoColumns)
				for _, row := range test.mock.expectRows {
					mockRows.AddRow(row...)
				}
				query.WillReturnRows(mockRows)
			}

			mateos, err := SelectAllMateos()

			assert.Equal(t, test.want.mateoMap, mateos)
			assert.Equal(t, test.want.err, err)
			assert.Nil(t, dbMock.ExpectationsWereMet())
		})
	}
}

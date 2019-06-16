package database

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func TestSelectAllMateos(t *testing.T) {
	type want struct {
		mateoMap []map[string]interface{}
		err      error
	}
	tests := []struct {
		name string
		mock Mock
		want
	}{
		{
			name: "Single Row Returned",
			mock: Mock{
				ExpectRows: []MockRow{
					GetValues(model.TestJohn, ""),
				},
			},
			want: want{
				mateoMap: []map[string]interface{}{
					GetMap(model.TestJohn, ""),
				},
				err: nil,
			},
		},
		{
			name: "Multiple Rows Returned",
			mock: Mock{
				ExpectRows: []MockRow{
					GetValues(model.TestJohn, ""),
					GetValues(model.TestAdam, ""),
				},
			},
			want: want{
				mateoMap: []map[string]interface{}{
					GetMap(model.TestJohn, ""),
					GetMap(model.TestAdam, ""),
				},
				err: nil,
			},
		},
		{
			name: "Connection Error",
			mock: Mock{
				ExpectErr: errors.New("connection error"),
			},
			want: want{
				mateoMap: nil,
				err:      errors.New("connection error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			dbMock := GetMockDB(st)
			query := dbMock.ExpectQuery(SelectAllMateosExpectedQuery)

			if test.mock.ExpectErr != nil {
				query.WillReturnError(test.mock.ExpectErr)
			} else {
				mockRows := sqlmock.NewRows(model.GetMateoColumns())
				for _, row := range test.mock.ExpectRows {
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

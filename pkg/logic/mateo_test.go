package logic

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/wexel-nath/meat-night/pkg/database"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func TestGetAllMateos(t *testing.T) {
	type want struct {
		mateos []model.Mateo
		err    error
	}

	tests := []struct {
		name string
		mock database.Mock
		sort string
		want
	}{
		{
			name: "Single Row Returned",
			mock: database.Mock{
				ExpectQuery: database.SelectAllMateosExpectedQuery,
				ExpectRows: []database.MockRow{
					database.GetValues(database.TestJohn, ""),
				},
			},
			sort: "new",
			want: want{
				mateos: []model.Mateo{
					database.TestJohn,
				},
				err: nil,
			},
		},
		{
			name: "Single Row Returned - Legacy",
			mock: database.Mock{
				ExpectQuery: database.SelectAllMateosLegacyExpectedQuery,
				ExpectRows: []database.MockRow{
					database.GetValues(database.TestBob, model.TypeLegacy),
				},
			},
			sort: model.TypeLegacy,
			want: want{
				mateos: []model.Mateo{
					database.TestBob,
				},
				err: nil,
			},
		},
		{
			name: "Multiple Rows Returned - Legacy",
			mock: database.Mock{
				ExpectQuery: database.SelectAllMateosLegacyExpectedQuery,
				ExpectRows: []database.MockRow{
					database.GetValues(database.TestBob, model.TypeLegacy),
					database.GetValues(database.TestDavid, model.TypeLegacy),
				},
			},
			sort: model.TypeLegacy,
			want: want{
				mateos: []model.Mateo{
					database.TestBob,
					database.TestDavid,
				},
				err: nil,
			},
		},
		{
			name: "Connection Error",
			mock: database.Mock{
				ExpectErr: errors.New("connection error"),
			},
			want: want{
				mateos: nil,
				err:    errors.New("connection error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			dbMock := database.GetMockDB(st)
			query := dbMock.ExpectQuery(test.mock.ExpectQuery)

			if test.mock.ExpectErr != nil {
				query.WillReturnError(test.mock.ExpectErr)
			} else {
				mockRows := sqlmock.NewRows(model.GetMateoColumns())
				for _, row := range test.mock.ExpectRows {
					mockRows.AddRow(row...)
				}
				query.WillReturnRows(mockRows)
			}

			mateos, err := GetAllMateos(test.sort)

			assert.Equal(t, test.want.mateos, mateos)
			assert.Equal(t, test.want.err, err)
			assert.Nil(t, dbMock.ExpectationsWereMet())
		})
	}
}

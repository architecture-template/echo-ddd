package dao

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/architecture-template/echo-ddd/domain/model"
	dbConfig "github.com/architecture-template/echo-ddd/config/db"
)

func TestExample_FindByExampleKey(t *testing.T) {
	testCases := []struct {
		name            string
		mockRows        *sqlmock.Rows
		mockError       error
		expectedExample *model.Example
		expectedError   error
	}{
		{
			name: "正常系: レコードが存在する場合",
			mockRows: sqlmock.NewRows([]string{"id", "example_key", "example_name", "created_at", "updated_at"}).
				AddRow(
					1, 
					"test_key",
					"test_name",
					time.Date(2023,time.January, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				),
			mockError: nil,
			expectedExample: &model.Example{
				ID:          1,
				ExampleKey:  "test_key",
				ExampleName: "test_name",
				CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name:            "正常系: レコードが存在しない場合",
			mockRows:        sqlmock.NewRows([]string{"id", "example_key", "example_name", "created_at", "updated_at"}),
			mockError:       nil,
			expectedExample: nil,
			expectedError:   gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			sqlHandler := &dbConfig.SqlHandler{
				ReadConn:  gormDB,
				WriteConn: gormDB,
			}

			repo := NewExampleDao(sqlHandler)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			example, err := repo.FindByExampleKey("test_key")
			assert.Equal(t, tc.expectedExample, example)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

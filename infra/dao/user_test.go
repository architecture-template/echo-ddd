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

func TestUser_FindByEmail(t *testing.T) {
	testCases := []struct {
		name          string
		mockRows      *sqlmock.Rows
		mockError     error
		expectedUser  *model.User
		expectedError error
	}{
		{
			name: "正常系: レコードが存在する場合",
			mockRows: sqlmock.NewRows([]string{"id", "user_key", "user_name", "email", "password", "token", "created_at", "updated_at"}).
				AddRow(
					1, 
					"test_key",
					"test_name",
					"test@test.com",
					"test_password",
					"test_token",
					time.Date(2023,time.January, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				),
			mockError: nil,
			expectedUser: &model.User{
				ID:        1,
				UserKey:   "test_key",
				UserName:  "test_name",
				Email:     "test@test.com",
				Password:  "test_password",
				Token:     "test_token",
				CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name:          "正常系: レコードが存在しない場合",
			mockRows:      sqlmock.NewRows([]string{"id", "user_key", "user_name", "email", "password", "token", "created_at", "updated_at"}),
			mockError:     nil,
			expectedUser:  nil,
			expectedError: gorm.ErrRecordNotFound,
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

			repo := NewUserDao(sqlHandler)
			mock.ExpectQuery("SELECT").WithArgs("test@test.com").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			user, err := repo.FindByEmail("test@test.com")
			assert.Equal(t, tc.expectedUser, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUser_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.User
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedUser     *model.User
        expectedError    error
    }{
        {
            name: "正常系",
            mockParam: &model.User{
                UserKey:  "test_key",
                UserName: "test_name",
                Email:    "test@test.com",
                Password: "test_password",
                Token:    "test_token",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedUser: &model.User{
                ID:        1,
                UserKey:   "test_key",
                UserName: "test_name",
                Email:     "test@test.com",
                Password:  "test_password",
                Token:     "test_token",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
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
			repo := NewUserDao(sqlHandler)

            mock.ExpectBegin()
            mock.ExpectExec("INSERT").
                WillReturnResult(sqlmock.NewResult(tc.mockLastInsertID, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            user, err := repo.Insert(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

            user.CreatedAt = time.Time{}
            user.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedUser, user)
        })
    }
}

func TestUser_Update(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.User
        mockRowsAffected int64
        mockLastUpdateID int64
        mockError        error
        expectedUser     *model.User
        expectedError    error
    }{
        {
			name: "正常系",
            mockParam: &model.User{
                UserKey:  "test_key",
                UserName: "test_name",
                Email:    "test@test.com",
                Password: "test_password",
                Token:    "test_token",
            },
            mockRowsAffected: 1,
            mockLastUpdateID: 1,
            mockError:        nil,
            expectedUser: &model.User{
                ID:        0,
                UserKey:   "test_key",
                UserName:  "test_name",
                Email:     "test@test.com",
                Password:  "test_password",
                Token:     "test_token",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
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
			repo := NewUserDao(sqlHandler)

            mock.ExpectBegin()
			mock.ExpectExec("UPDATE").
                WillReturnResult(sqlmock.NewResult(tc.mockLastUpdateID, tc.mockRowsAffected)).
				WillReturnError(tc.mockError)
            mock.ExpectCommit()

            user, err := repo.Update(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

			user.ID = 0
            user.CreatedAt = time.Time{}
            user.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedUser, user)
        })
    }
}

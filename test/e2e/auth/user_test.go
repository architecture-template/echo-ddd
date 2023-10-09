package e2e

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"

	"github.com/architecture-template/echo-ddd/auth/presentation/parameter"
	"github.com/architecture-template/echo-ddd/auth/presentation/output"
	"github.com/architecture-template/echo-ddd/auth/presentation/response"
)

func TestUser_RegisterUser(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		body         *parameter.RegisterUser
		expectedCode int
		expectedKey  string
	}{
		{
			name: "正常系: ユーザー登録",
			body: &parameter.RegisterUser{
				UserName: "test_name",
				Email:    "test@test.com",
				Password: "test_password",
			},
			expectedCode: http.StatusOK,
			expectedKey:  "test_key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatalf("JSON encoding error: %v", err)
				return
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/register_user", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Fatalf("Expected status code %v, but got %v", tc.expectedCode, resp.StatusCode)
			}

			if tc.expectedCode == http.StatusOK {
				actual := &response.Success{
					Items: &output.User{},
				}
				expect := &response.Success{
					Types: "register_user",
					Status: 200,
					Items: &output.User{
						UserKey:  "test_key",
						UserName: "test_name",
						Email:    "test@test.com",
						Token:    "nil",
						Message:  "user completed",						
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if userRegister, ok := actual.Items.(*output.User); ok {
					userRegister.UserKey = "test_key"
				}

				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestUser_LoginUser(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
		"sql/user/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		body         *parameter.LoginUser
		expectedCode int
	}{
		{
			name: "正常系: ユーザーログイン",
			body: &parameter.LoginUser{
				Email:     "test@test.com",
				Password:  "test_password",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatalf("JSON encoding error: %v", err)
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/login_user", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Fatalf("Expected status code %v, but got %v", tc.expectedCode, resp.StatusCode)
			}

			if tc.expectedCode == http.StatusOK {
				actual := &response.Success{
					Items: &output.User{},
				}
				expect := &response.Success{
					Types: "login_user",
					Status: 200,
					Items: &output.User{
						UserKey:  "test_key",
						UserName: "test_name",
						Email:    "test@test.com",
						Token:    "test_token",
						Message:  "user completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if user, ok := actual.Items.(*output.User); ok {
					user.Token = "test_token"
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

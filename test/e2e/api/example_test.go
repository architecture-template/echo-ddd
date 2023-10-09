package e2e

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"

	"github.com/architecture-template/echo-ddd/api/presentation/output"
	"github.com/architecture-template/echo-ddd/api/presentation/response"
)

func TestExample_GetExample(t *testing.T) {
	files := []File{
		"sql/example/example_table.sql",
		"sql/example/example_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		exampleKey   string
		exampleName  string
		expectedCode int
	}{
		{
			name:        "正常系: データが存在する場合",
			exampleKey:  "test_key",
			exampleName: "test_name",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/example/%s/get_example", os.Getenv("TEST_API_URL"), tc.exampleKey), nil)
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
					Items: &output.Example{},
				}
				expect := &response.Success{
					Types:  "get_example",
					Status: 200,
					Items: &output.Example{
						ExampleKey:  tc.exampleKey,
						ExampleName: tc.exampleName,
						Message:     "get example completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

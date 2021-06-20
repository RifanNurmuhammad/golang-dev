package omdb

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/rifannurmuhammad/02-movie-service/component"
	"github.com/stretchr/testify/assert"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mockhost = "http://omdb.mock"
	mockkey  = "abc"
)

// MockHTTP for mocking http server
func MockHTTP(method, url string, responseBodyData interface{}) {
	httpmock.RegisterResponder(method, url, func(*http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(http.StatusOK, responseBodyData)
		if err != nil {
			return httpmock.NewStringResponse(500, "error mock http"), nil
		}
		return resp, nil
	})
}

func TestNewBarracudaService(t *testing.T) {
	type env struct {
		key   string
		value string
	}

	tests := []struct {
		name    string
		env     []env
		wantErr bool
	}{
		{
			name: "Testcase #1: Test success create new service",
			env: []env{
				{key: "OMDB_URL", value: mockhost},
				{key: "OMDB_KEY", value: mockkey},
			},
		},
		{
			name:    "Testcase #2: Test error create new service (invalid env variable)",
			env:     []env{},
			wantErr: true,
		},
		{
			name: "Testcase #3: Test error create new service (invalid env variable)",
			env: []env{
				{key: "OMDB_KEY", value: mockkey},
			},
			wantErr: true,
		},
		{
			name: "Testcase #4: Test error create new service (url error)",
			env: []env{
				{key: "OMDB_URL", value: "<:::>"},
				{key: "OMDB_KEY", value: mockkey},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, env := range tt.env {
				os.Setenv(env.key, env.value)
			}

			db, _, _ := sqlMock.New()
			defer db.Close()
			gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
			_, err = NewOmdbService(&component.Mysql{Write: gdb, Read: gdb})
			if (err != nil) != tt.wantErr {
				t.Errorf("\x1b[31;NewOmdbService() error = %v, wantErr %v\x1b[0m", err, tt.wantErr)
			}

			for _, env := range tt.env {
				os.Unsetenv(env.key)
			}
		})
	}
}

func Test_omdb_Find(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	os.Setenv("OMDB_KEY", mockkey)
	os.Setenv("OMDB_URL", mockhost)

	tests := []struct {
		name             string
		expectedResponse ResponseOmdb
		param            string
		wantErr          bool
	}{
		{
			name:             "Testcase #1: Test success",
			expectedResponse: ResponseOmdb{Response: "True"},
			param:            "batman",
			wantErr:          false,
		},
		{
			name:             "Testcase #2: Test negative",
			expectedResponse: ResponseOmdb{Response: "False"},
			param:            "batman",
			wantErr:          true,
		},
		{
			name:             "Testcase #2: Test negative",
			expectedResponse: ResponseOmdb{Response: "False"},
			param:            "",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", mockhost, mockkey, "batman", 1)

		MockHTTP(http.MethodGet, url, tt.expectedResponse)

		t.Run(tt.name, func(t *testing.T) {
			db, _, _ := sqlMock.New()
			defer db.Close()
			gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
			service, err := NewOmdbService(&component.Mysql{Write: gdb, Read: gdb})
			if err != nil {
				assert.Error(t, err)
			}

			_, err = service.Find(tt.param, 0)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func Test_omdb_FindDetail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	os.Setenv("OMDB_KEY", mockkey)
	os.Setenv("OMDB_URL", mockhost)
	tests := []struct {
		name             string
		expectedResponse ResponseDetailOmdb
		param            string
		wantErr          bool
	}{
		{
			name:             "Testcase #1: Test success",
			expectedResponse: ResponseDetailOmdb{Response: "True"},
			param:            "id",
			wantErr:          false,
		},
		{
			name:             "Testcase #2: Test negative",
			expectedResponse: ResponseDetailOmdb{Response: "False"},
			param:            "id",
			wantErr:          true,
		},
		{
			name:             "Testcase #2: Test negative",
			expectedResponse: ResponseDetailOmdb{Response: "False"},
			param:            "",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		url := fmt.Sprintf("%s?apikey=%s&i=%s", mockhost, mockkey, "id")

		MockHTTP(http.MethodGet, url, tt.expectedResponse)

		t.Run(tt.name, func(t *testing.T) {
			db, _, _ := sqlMock.New()
			defer db.Close()
			gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
			service, err := NewOmdbService(&component.Mysql{Write: gdb, Read: gdb})
			if err != nil {
				assert.Error(t, err)
			}

			_, err = service.FindDetail(tt.param)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

package mysql

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/rifannurmuhammad/02-movie-service/component"
	"github.com/rifannurmuhammad/02-movie-service/src/log/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type testCase struct {
	name, expectQuery string
	wantError         bool
	wantServiceError  error
}

func TestNewLogRequestRepo(t *testing.T) {
	db, _, _ := sqlMock.New()
	defer db.Close()
	gdb, _ := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	NewLogActivitiesRepo(&component.Mysql{Write: gdb, Read: gdb})
}

func Test_log_insert(t *testing.T) {
	// set table name
	db, mock, _ := sqlMock.New()
	defer db.Close()
	gdb, _ := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	repository := NewLogActivitiesRepo(&component.Mysql{Write: gdb, Read: gdb})

	tests := []testCase{
		{
			name:      "Testcase #1: Positive",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			// set mock based on query
			mocks := mock.ExpectQuery(regexp.QuoteMeta(tt.expectQuery))

			if !tt.wantError {
				mocks = mocks.WillReturnRows(sqlMock.NewRows([]string{"id", "searchCall", "createdAt"}).AddRow("1", "abc", time.Now()))
			} else {
				mocks = mocks.WillReturnError(fmt.Errorf("error"))
			}

			err := <-repository.Insert(context.Background(), &domain.LogActivities{})
			if !tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

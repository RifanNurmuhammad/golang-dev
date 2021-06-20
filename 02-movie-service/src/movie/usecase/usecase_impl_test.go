package usecase

import (
	"fmt"
	"testing"

	"github.com/rifannurmuhammad/02-movie-service/service"
	"github.com/rifannurmuhammad/02-movie-service/service/omdb"
	omdbServiceMock "github.com/rifannurmuhammad/02-movie-service/service/omdb/mocks"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name             string
	wantError        bool
	wantServiceError error
}

var tests = []testCase{
	{
		name:      "Testcase #1: Positive",
		wantError: false,
	},
	{
		name:             "Testcase #2: Negative",
		wantError:        true,
		wantServiceError: fmt.Errorf("error"),
	},
}

func Test_NewMovieUsecase(t *testing.T) {
	usecase := NewMovieUsecase(new(service.Service))
	assert.NotNil(t, usecase)
}

func Test_usecase_GetMovies(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockService := new(omdbServiceMock.Omdb)
			mockService.On("Find", mock.Anything, mock.Anything).Return(omdb.ResponseOmdb{Search: []omdb.SearchDetail{{Title: "abc"}}}, tt.wantServiceError)
			uc := new(MovieUseCaseImpl)

			uc.service = &service.Service{Omdb: mockService}

			_, _, err := uc.GetMovies(domain.Parameters{})
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_usecase_GetMovieDetail(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockService := new(omdbServiceMock.Omdb)
			mockService.On("FindDetail", mock.Anything).Return(omdb.ResponseDetailOmdb{Ratings: []omdb.RatingDetail{{Source: "abc"}}}, tt.wantServiceError)
			uc := new(MovieUseCaseImpl)

			uc.service = &service.Service{Omdb: mockService}

			_, err := uc.GetMovieDetail("1")
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

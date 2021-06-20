package delivery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/rifannurmuhammad/02-movie-service/shared"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
	usecaseMock "github.com/rifannurmuhammad/02-movie-service/src/movie/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name             string
	wantUsecaseError error
	wantErr          bool
	wantStatusCode   int
}

var testHttp = []testCase{
	{
		name:           "Testcase #1: Positive",
		wantStatusCode: http.StatusOK,
	},
	{
		name:             "Testcase #2: Negative",
		wantStatusCode:   http.StatusBadRequest,
		wantUsecaseError: fmt.Errorf("error pq"),
	},
}

func TestHandler_Mount(t *testing.T) {
	e := echo.New()
	handler := NewHTTPHandler(new(usecaseMock.MovieUseCase))
	handler.Mount(e.Group("/"))
	handler.Mount(e.Group("/:id"))
}

func TestHandler_getMovies(t *testing.T) {
	for _, tt := range testHttp {
		t.Run(tt.name, func(t *testing.T) {

			mockUsecase := new(usecaseMock.MovieUseCase)
			mockUsecase.On("GetMovies", mock.Anything).Return([]domain.Movies{}, shared.Meta{}, tt.wantUsecaseError)

			e := echo.New()
			uri := fmt.Sprintf("/api/movie?searchword=%s&pagination=%s", "batman", "1")
			req, err := http.NewRequest(echo.GET, uri, nil)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := NewHTTPHandler(mockUsecase)
			handler.GetMovies(c)

			assert.Equal(t, tt.wantStatusCode, rec.Code)
		})
	}
}

func TestHandler_getMovieDetail(t *testing.T) {
	for _, tt := range testHttp {
		t.Run(tt.name, func(t *testing.T) {

			mockUsecase := new(usecaseMock.MovieUseCase)
			mockUsecase.On("GetMovieDetail", mock.Anything).Return(domain.MovieDetail{}, tt.wantUsecaseError)

			e := echo.New()
			uri := fmt.Sprintf("/api/movie/%s", "1")
			req, err := http.NewRequest(echo.GET, uri, nil)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := NewHTTPHandler(mockUsecase)
			handler.GetMovieDetail(c)

			assert.Equal(t, tt.wantStatusCode, rec.Code)
		})
	}
}

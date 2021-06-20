package delivery

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/rifannurmuhammad/02-movie-service/grpc/proto"
	"github.com/rifannurmuhammad/02-movie-service/shared"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
	usecaseMock "github.com/rifannurmuhammad/02-movie-service/src/movie/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var test = []testCase{
	{
		name:    "Testcase #1: Positive",
		wantErr: false,
	},
	{
		name:             "Testcase #2: Negative",
		wantErr:          true,
		wantUsecaseError: fmt.Errorf("error pq"),
	},
}

func TestGrpcHandler_getMovie(t *testing.T) {

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {

			mockUsecase := new(usecaseMock.MovieUseCase)
			mockUsecase.On("GetMovies", mock.Anything).Return([]domain.Movies{{Title: "abc"}}, shared.Meta{}, tt.wantUsecaseError)

			handler := NewGrpcHandler(mockUsecase)
			_, err := handler.GetMovie(context.Background(), &pb.MovieQuery{})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGrpcHandler_getMovieDetail(t *testing.T) {
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {

			mockUsecase := new(usecaseMock.MovieUseCase)
			mockUsecase.On("GetMovieDetail", mock.Anything).Return(domain.MovieDetail{}, tt.wantUsecaseError)

			handler := NewGrpcHandler(mockUsecase)
			_, err := handler.GetMovieDetail(context.Background(), &pb.MovieDetailQuery{})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

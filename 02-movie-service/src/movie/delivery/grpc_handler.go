package delivery

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rifannurmuhammad/02-movie-service/grpc/proto"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/usecase"
)

// GrpcHandler struct
type GrpcHandler struct {
	MovieUseCase usecase.MovieUseCase
}

// NewGrpcHandler function, for intialize GrpcHandler object
func NewGrpcHandler(MovieUseCase usecase.MovieUseCase) *GrpcHandler {
	return &GrpcHandler{MovieUseCase}
}

// GetMovie function, implementation function from Movie Protobuf
func (h *GrpcHandler) GetMovie(ctx context.Context, arg *pb.MovieQuery) (*pb.ResponseMovieList, error) {

	pagination, _ := strconv.Atoi(arg.Pagination)
	params := domain.Parameters{
		Pagination: pagination,
		SearchWord: arg.SearchWord,
	}

	movieResult, _, err := h.MovieUseCase.GetMovies(params)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	responseMovieList := []*pb.MovieList{}
	for _, data := range movieResult {
		movieList := &pb.MovieList{
			Title:  data.Title,
			Year:   data.Year,
			ImdbID: data.ImdbID,
			Type:   data.Type,
			Poster: data.Poster,
		}
		responseMovieList = append(responseMovieList, movieList)
	}

	result := &pb.ResponseMovieList{}
	result.Movie = responseMovieList

	return result, nil
}

// GetMovieDetail function for getting detail of movie
func (h *GrpcHandler) GetMovieDetail(ctx context.Context, id *pb.MovieDetailQuery) (*pb.MovieDetail, error) {
	movieResult, err := h.MovieUseCase.GetMovieDetail(id.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := &pb.MovieDetail{
		Title:      movieResult.Title,
		Year:       movieResult.Year,
		Rated:      movieResult.Rated,
		Released:   movieResult.Released,
		Runtime:    movieResult.Runtime,
		Genre:      movieResult.Genre,
		Director:   movieResult.Director,
		Writer:     movieResult.Writer,
		Actors:     movieResult.Actors,
		Plot:       movieResult.Plot,
		Language:   movieResult.Language,
		Country:    movieResult.Country,
		Awards:     movieResult.Awards,
		Poster:     movieResult.Poster,
		MetaScore:  movieResult.MetaScore,
		ImdbRating: movieResult.ImdbRating,
		ImdbVotes:  movieResult.ImdbVotes,
		ImdbID:     movieResult.ImdbID,
		Type:       movieResult.Type,
		DVD:        movieResult.DVD,
		BoxOffice:  movieResult.BoxOffice,
		Production: movieResult.Production,
		Website:    movieResult.Website,
	}
	return result, nil
}

package usecase

import (
	"github.com/rifannurmuhammad/02-movie-service/shared"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
)

// MovieUseCase abstraction
type MovieUseCase interface {
	GetMovies(param domain.Parameters) (data []domain.Movies, meta shared.Meta, err error)
	GetMovieDetail(id string) (data domain.MovieDetail, err error)
}

package usecase

import (
	"strconv"

	"github.com/rifannurmuhammad/02-movie-service/service"
	"github.com/rifannurmuhammad/02-movie-service/shared"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
)

// MovieUseCaseImpl data structure
type MovieUseCaseImpl struct {
	service *service.Service
}

// NewMovieUsecase constructor
func NewMovieUsecase(service *service.Service) MovieUseCase {
	return &MovieUseCaseImpl{
		service: service,
	}
}

// GetMovies function for get list of movie
func (movie *MovieUseCaseImpl) GetMovies(param domain.Parameters) (data []domain.Movies, meta shared.Meta, err error) {

	result, err := movie.service.Omdb.Find(param.SearchWord, param.Pagination)
	if err != nil {
		return
	}
	for _, movieData := range result.Search {
		data = append(data, domain.Movies{
			Title:  movieData.Title,
			Year:   movieData.Year,
			ImdbID: movieData.ImdbID,
			Type:   movieData.Type,
			Poster: movieData.Poster,
		})
	}

	totalResult, _ := strconv.Atoi(result.TotalResults)
	meta = shared.NewMeta(param.Pagination, totalResult)
	return
}

// GetMovies function for get list of movie
func (movie *MovieUseCaseImpl) GetMovieDetail(id string) (data domain.MovieDetail, err error) {

	result, err := movie.service.Omdb.FindDetail(id)
	if err != nil {
		return
	}

	data.Title = result.Title
	data.Year = result.Year
	data.Rated = result.Rated
	data.Released = result.Released
	data.Runtime = result.Runtime
	data.Genre = result.Genre
	data.Director = result.Director
	data.Writer = result.Writer
	data.Actors = result.Actors
	data.Plot = result.Plot
	data.Language = result.Language
	data.Country = result.Country
	data.Awards = result.Awards
	data.Poster = result.Poster
	data.MetaScore = result.MetaScore
	data.ImdbRating = result.ImdbRating
	data.ImdbVotes = result.ImdbVotes
	data.ImdbID = result.ImdbID
	data.Type = result.Type
	data.DVD = result.DVD
	data.BoxOffice = result.BoxOffice
	data.Production = result.Production
	data.Website = result.Website
	ratings := []domain.RatingDetail{}
	for _, rating := range result.Ratings {
		ratings = append(ratings,
			domain.RatingDetail{Source: rating.Source, Value: rating.Value},
		)
	}
	data.Ratings = ratings
	return
}

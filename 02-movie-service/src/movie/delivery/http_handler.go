package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rifannurmuhammad/02-movie-service/helper"
	"github.com/rifannurmuhammad/02-movie-service/shared"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/domain"
	"github.com/rifannurmuhammad/02-movie-service/src/movie/usecase"
)

// HTTPMovieHandler structure
type HTTPMovieHandler struct {
	MovieUseCase usecase.MovieUseCase
}

// NewHTTPHandler function for initialise *HTTPMovieHandler
func NewHTTPHandler(movieUseCase usecase.MovieUseCase) *HTTPMovieHandler {
	return &HTTPMovieHandler{MovieUseCase: movieUseCase}
}

// Mount function for mounting routes
func (h *HTTPMovieHandler) Mount(group *echo.Group) {
	group.GET("", h.GetMovies)
	group.GET("/:id", h.GetMovieDetail)
}

// GetMovies function for getting list of movie
func (h *HTTPMovieHandler) GetMovies(c echo.Context) error {
	pagination, _ := strconv.Atoi(c.QueryParam("pagination"))
	params := domain.Parameters{
		Pagination: pagination,
		SearchWord: c.QueryParam("searchword"),
	}

	movieResult, meta, err := h.MovieUseCase.GetMovies(params)
	if err != nil {
		return shared.NewHTTPResponse(http.StatusBadRequest, err.Error(), make(helper.EmptySlice, 0)).JSON(c)
	}

	return shared.NewHTTPResponse(http.StatusOK, "Success Get All Movies", movieResult, meta).JSON(c)
}

// GetMovieDetail function for getting detail of movie
func (h *HTTPMovieHandler) GetMovieDetail(c echo.Context) error {
	id := c.Param("id")

	movieResult, err := h.MovieUseCase.GetMovieDetail(id)
	if err != nil {
		return shared.NewHTTPResponse(http.StatusBadRequest, err.Error(), make(helper.EmptySlice, 0)).JSON(c)
	}

	return shared.NewHTTPResponse(http.StatusOK, "Success Get Detail Movie", movieResult).JSON(c)
}

package omdb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/rifannurmuhammad/02-movie-service/shared"
)

// Omdb service
type Omdb interface {
	Find(search string, page int) (ResponseOmdb, error)
	FindDetail(id string) (ResponseDetailOmdb, error)
}

// omdbService data structure for baracuda receiver
type omdbService struct {
	baseURL *url.URL
	key     string
}

// NewOmdbService function for initializing omdb service
func NewOmdbService() (Omdb, error) {
	var (
		err         error
		ok          bool
		omdbservice = new(omdbService)
	)

	omdbservice.key, ok = os.LookupEnv("OMDB_KEY")
	if !ok {
		err := errors.New("you need to specify BARRACUDA_BASIC_AUTH in the environment variable")
		return nil, err
	}

	baseURL, ok := os.LookupEnv("OMDB_URL")
	if !ok {
		err := errors.New("you need to specify BARRACUDA_BASE_URL in the environment variable")
		return nil, err
	}

	omdbservice.baseURL, err = url.Parse(baseURL)
	if err != nil {
		err = errors.New("url is invalid")
		return nil, err
	}

	return omdbservice, nil
}

// Find is method from OmdbService for get movie
func (omdb *omdbService) Find(search string, page int) (ResponseOmdb, error) {

	var responseOmdb ResponseOmdb
	if page <= 0 {
		page = 1
	}
	uri := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", omdb.baseURL.String(), omdb.key, search, page)
	err := shared.HTTPRequest(context.Background(), http.MethodGet, uri, nil, &responseOmdb, nil)
	if err != nil {
		return responseOmdb, err
	}

	if responseOmdb.Response == "False" {
		return responseOmdb, errors.New(responseOmdb.Error)
	}

	return responseOmdb, nil
}

// FindDetail is method from OmdbService for get movie by id
func (omdb *omdbService) FindDetail(id string) (ResponseDetailOmdb, error) {
	var responseDetail ResponseDetailOmdb
	uri := fmt.Sprintf("%s?apikey=%s&i=%s", omdb.baseURL.String(), omdb.key, id)
	err := shared.HTTPRequest(context.Background(), http.MethodGet, uri, nil, &responseDetail, nil)
	if err != nil {
		return responseDetail, err
	}

	if responseDetail.Response == "False" {
		return responseDetail, errors.New(responseDetail.Error)
	}

	return responseDetail, nil
}
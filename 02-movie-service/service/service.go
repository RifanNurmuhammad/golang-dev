package service

import (
	"fmt"

	"github.com/rifannurmuhammad/02-movie-service/service/omdb"
)

// ResultService model
type ResultService struct {
	Result interface{}
	Error  error
}

// Service for common service
type Service struct {
	Omdb omdb.Omdb
}

// NewService construct common service
func NewService() *Service {
	var err error
	srv := new(Service)

	srv.Omdb, err = omdb.NewOmdbService()
	if err != nil {
		err = fmt.Errorf("Panic. Cannot load service, please check environment. Error: %v", err)
	}
	return srv
}

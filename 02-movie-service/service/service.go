package service

import (
	"fmt"

	"github.com/rifannurmuhammad/02-movie-service/component"
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
func NewService(repoDb *component.Mysql) *Service {
	var err error
	srv := new(Service)

	srv.Omdb, err = omdb.NewOmdbService(repoDb)
	if err != nil {
		err = fmt.Errorf("Panic. Cannot load service, please check environment. Error: %v", err)
	}
	return srv
}

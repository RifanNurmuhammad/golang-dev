package main

import (
	"github.com/rifannurmuhammad/02-movie-service/component"
	mysqlConfig "github.com/rifannurmuhammad/02-movie-service/config/mysql"
	movieUsecasePackage "github.com/rifannurmuhammad/02-movie-service/src/movie/usecase"

	servicePackage "github.com/rifannurmuhammad/02-movie-service/service"
)

// Movie main service structure
type Service struct {
	MovieUseCase movieUsecasePackage.MovieUseCase
	Service      *servicePackage.Service
}

// MakeHandler function for initializing service
func MakeHandler() *Service {
	dbConnection := &component.Mysql{Write: mysqlConfig.LoadMysqlDB(), Read: mysqlConfig.LoadMysqlDB()}
	serviceModule := servicePackage.NewService(dbConnection)
	movieUseCase := movieUsecasePackage.NewMovieUsecase(serviceModule)

	service := new(Service)
	service.MovieUseCase = movieUseCase
	service.Service = serviceModule
	return service
}

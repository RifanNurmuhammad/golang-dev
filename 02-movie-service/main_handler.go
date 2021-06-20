package main

import (
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
	serviceModule := servicePackage.NewService()
	movieUseCase := movieUsecasePackage.NewMovieUsecase(serviceModule)

	service := new(Service)
	service.MovieUseCase = movieUseCase
	service.Service = serviceModule
	return service
}

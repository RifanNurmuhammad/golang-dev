package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo"
	movieDelivery "github.com/rifannurmuhammad/02-movie-service/src/movie/delivery"
)

// HTTPDefaultPort , default port for HTTP Server
const HTTPDefaultPort = 8080

// HTTPServerMain entry function for initializing http server
func (s Service) HTTPServerMain() {

	e := echo.New()

	movieHandler := movieDelivery.NewHTTPHandler(s.MovieUseCase)
	movieGroup := e.Group("/api/movie")
	movieHandler.Mount(movieGroup)

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = HTTPDefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = HTTPDefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}

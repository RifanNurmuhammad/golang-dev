package main

import (
	"fmt"
	"os"

	"github.com/rifannurmuhammad/02-movie-service/component"
	mysqlConfig "github.com/rifannurmuhammad/02-movie-service/config/mysql"

	rpc "github.com/rifannurmuhammad/02-movie-service/grpc/server"
	servicePackage "github.com/rifannurmuhammad/02-movie-service/service"
	movieDeliveryPackage "github.com/rifannurmuhammad/02-movie-service/src/movie/delivery"
	movieUsecasePackage "github.com/rifannurmuhammad/02-movie-service/src/movie/usecase"
)

// GRPCDefaultPort default port for GRPC
const GRPCDefaultPort = 6565

// GRPCMain entry function for initializing GRPC Server
func (s *Service) GRPCMain() {
	dbConnection := &component.Mysql{Write: mysqlConfig.LoadMysqlDB(), Read: mysqlConfig.LoadMysqlDB()}
	serviceModule := servicePackage.NewService(dbConnection)
	movieUseCase := movieUsecasePackage.NewMovieUsecase(serviceModule)
	movieGrpcHandler := movieDeliveryPackage.NewGrpcHandler(movieUseCase)

	fmt.Printf("GRPC Server Running on port = %d ", GRPCDefaultPort)

	grpcServer, err := rpc.NewGrpcServer(movieGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(GRPCDefaultPort)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}
}

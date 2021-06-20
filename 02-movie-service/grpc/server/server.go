package servers

import (
	"fmt"
	"net"

	pb "github.com/rifannurmuhammad/02-movie-service/grpc/proto"
	movieDelivery "github.com/rifannurmuhammad/02-movie-service/src/movie/delivery"

	"google.golang.org/grpc"
)

//Server data structure, grpc server model
type Server struct {
	movieGrpcHandler *movieDelivery.GrpcHandler
	pb.UnimplementedMovieServiceServer
}

// Serve insecure server/ no server side encryption
func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf("localhost:%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	server := &Server{}
	//Register all sub server here
	pb.RegisterMovieServiceServer(grpcServer, server)
	//end register server

	err = grpcServer.Serve(l)

	if err != nil {
		return err
	}

	return nil

}

//NewGrpcServer function, return: Pointer GRPC Server, or error otherwise
func NewGrpcServer(movieGrpcHandler *movieDelivery.GrpcHandler) (*Server, error) {
	return &Server{
		movieGrpcHandler: movieGrpcHandler,
	}, nil

}

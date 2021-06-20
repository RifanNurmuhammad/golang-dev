package main

import (
	"context"
	"log"

	pb "github.com/rifannurmuhammad/02-movie-service/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":6565", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewMovieServiceClient(conn)

	response, err := client.GetMovie(context.Background(), &pb.MovieQuery{Pagination: "1", SearchWord: "Batman"})
	log.Printf("Error from server: %v \n", err)
	log.Printf("Response from server: %v \n", response)
}

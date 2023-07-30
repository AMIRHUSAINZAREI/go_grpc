package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/AMIRHUSAINZAREI/go_grpc_sample/pkg"
	"github.com/AMIRHUSAINZAREI/go_grpc_sample/pkg/calc"
	pb "github.com/AMIRHUSAINZAREI/go_grpc_sample/proto/calc"
	"google.golang.org/grpc"
)

type calcServer struct {
	pb.UnimplementedCalcServer
}

func (sc *calcServer) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := calc.Add(int(in.A), int(in.B))
	return &pb.Response{Result: float32(result)}, nil
}

func (sc *calcServer) Sub(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := calc.Sub(int(in.A), int(in.B))
	return &pb.Response{Result: float32(result)}, nil
}

func (sc *calcServer) Mul(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := calc.Mul(int(in.A), int(in.B))
	return &pb.Response{Result: float32(result)}, nil
}

func (sc *calcServer) Div(ctx context.Context, in *pb.Request) (*pb.ResponseWithError, error) {
	result, err := calc.Div(int(in.A), int(in.B))
	return &pb.ResponseWithError{Result: float32(result), Error: err.Error()}, nil
}

func main() {
	// Load port number from .env file
	port, err := pkg.GetEnv("CALC_GRPC_SERVER_PORT")
	if err != nil {
		log.Fatalf("Failed to load CALC_GRPC_SERVER_PORT environment variable")
	}

	// Create a tcp connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	// Register CalcServer
	server := grpc.NewServer()
	pb.RegisterCalcServer(server, &calcServer{})

	// Start listening
	fmt.Println("Server started. Listening on port", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

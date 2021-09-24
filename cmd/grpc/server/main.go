package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ldenholm/porus/internal/grpc/impl"
	"github.com/ldenholm/porus/pb"
	"google.golang.org/grpc"
)

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	serv := impl.AccountServiceImpl{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, &serv)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

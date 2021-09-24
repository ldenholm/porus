// package main

// import (
// 	"log"
// 	"net"

// 	"github.com/ldenholm/porus/pb"
// 	"google.golang.org/grpc"
// )

// // gRPC API for Account Service

// const (
// 	port     = ":50051"
// 	clientID = "account-api"
// )

// // Types
// type AccountServiceServer struct {
// 	GetBalance __.AccountServiceServer.GetBalance()
// }

// // Method Implementation
// //func (s *server) GetBalance(BalanceRequest) returns (BalanceResponse) {}

// // Server
// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	// Creates gRPC server
// 	s := grpc.NewServer()
// 	srv := __.AccountServiceServer.GetBalance()
// 	//__.RegisterAccountServiceServer(s, )
// 	s.Serve(lis)
// }

// // Connect to NATS
// // err = comp.ConnectToNATSStreaming(
// // 	clusterID,
// // 	stan.NatsURL(stan.DefaultNatsURL),
// // )
// // if err != nil {
// // 	log.Fatal(err)
// // }

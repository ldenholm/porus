package main

import (
	"log"

	"github.com/ldenholm/porus/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Hard coded example of a client (not currently used).

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := pb.NewAccountServiceClient(conn)

	response, err := c.GetBalance(context.Background(), &pb.BalanceRequest{UserId: 1})

	if err != nil {
		log.Fatalf("Error when calling GetBalance: %s", err)
	}
	log.Printf("Response from server: %x", response.Balance)
}

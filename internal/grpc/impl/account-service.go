package impl

import (
	"context"
	"log"

	"github.com/ldenholm/porus/pb"
)

//AccountServiceImpl is a implementation of RepositoryService Grpc Service.

type AccountServiceImpl struct {
}

//AccountServiceImpl returns the pointer to the implementation.
func NewAccountServiceImpl() *AccountServiceImpl {
	return &AccountServiceImpl{}
}

func (serviceImpl *AccountServiceImpl) GetBalance(ctx context.Context, in *pb.BalanceRequest) (*pb.BalanceResponse, error) {
	log.Println("Received get balance request")
	// lookup user in db and return balance (in.UserId)
	if in.UserId != 1 {
		return nil, ctx.Err()
	}

	return &pb.BalanceResponse{
		Balance: 0,
	}, nil
}

func (serviceImpl *AccountServiceImpl) AddFunds(ctx context.Context, in *pb.AddFundsRequest) (*pb.AddFundsResponse, error) {
	log.Println("Received add funds request")
	return &pb.AddFundsResponse{
		Total: 700,
	}, nil
}

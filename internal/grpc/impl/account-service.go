package impl

import (
	"context"
	"log"
)

//AccountServiceImpl is a implementation of RepositoryService Grpc Service.
type AccountServiceImpl struct {
}

//AccountServiceImpl returns the pointer to the implementation.
func NewAccountServiceImpl() *AccountServiceImpl {
	return &AccountServiceImpl{}
}

func (serviceImpl *AccountServiceImpl) GetBalance(ctx context.Context, in *BalanceRequest) (*BalanceResponse, error) {
	log.Println("Received get balance request")
	// lookup user in db and return balance (in.UserId)
	return &BalanceResponse{
		Balance: 0,
	}, nil
}

func (serviceImpl *AccountServiceImpl) AddFunds(ctx context.Context, in *AddFundsRequest) (*AddFundsResponse, error) {
	log.Println("Received add funds request")
	return &AddFundsResponse{
		Total: 700,
	}, nil
}

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

func (serviceImpl *AccountServiceImpl) GetBalance(ctx context.Context, userId int32) float32 {
	log.Println("Received get balance request")
	// lookup user in db and return balance
	return 32
}

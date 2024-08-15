package base

import (
	"context"

	"github.com/coderajay94/microservice1/api"
	"github.com/coderajay94/microservice1/model"
	"go.uber.org/zap"
)

//type Middleware func(s Service) Service

type Service interface {
	GetAccountDetails(ctx context.Context, req model.UserRequest) (model.UserResponse, error)
}

type baseService struct{
	memory api.MemoryUserAccounts
	logger *zap.Logger
	queryLimit int
}

func NewService(logger *zap.Logger, queryLimit int) Service{
       return baseService{
		memory: api.InitMemoryUserAccounts(),
		logger: logger,
		queryLimit: queryLimit,
	   }
}

func (b baseService)GetAccountDetails(_ context.Context, req model.UserRequest) (model.UserResponse, error) {
	return b.memory.GetAccountDetails(req)
}

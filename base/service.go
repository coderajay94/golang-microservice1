package base

import (
	"context"
	"fmt"

	"github.com/coderajay94/microservice1/db"
	"github.com/coderajay94/microservice1/model"
	"go.uber.org/zap"
)

type Middleware func(s Service) Service

type Service interface {
	GetAccountDetails(ctx context.Context, req model.UserRequestDB) (model.UserResponseDB, error)
	SaveAccountDetails(ctx context.Context, req model.UserResponseDB)(model.SaveResponseDB, error)
}

type baseService struct{
	//memory api.MemoryUserAccounts
	logger *zap.Logger
	mongoDB db.MongoDatabase
	queryLimit int
}

func NewService(logger *zap.Logger, mongoDB db.MongoDatabase, queryLimit int) Service{
       return baseService{
		//memory: api.InitMemoryUserAccounts(),
		logger: logger,
		mongoDB: mongoDB,
		queryLimit: queryLimit,
	   }
}

func (b baseService)GetAccountDetails(_ context.Context, req model.UserRequestDB) (model.UserResponseDB, error) {
	fmt.Println("calling from GetAccountDetails service")
	return b.mongoDB.GetAccountDetails(req)
}

func(b baseService)SaveAccountDetails(_ context.Context, req model.UserResponseDB)(model.SaveResponseDB, error){
	fmt.Println("calling from SaveAccountDetails service")
	return b.mongoDB.SaveAccountDetails(req)
}
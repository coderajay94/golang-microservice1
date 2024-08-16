package base

import (
	"context"
	"time"

	"github.com/coderajay94/microservice1/model"
	"go.uber.org/zap"
)

//how we're returning this Middlware
func NewLoggingMiddleware(logger *zap.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next Service
	logger *zap.Logger
}

func (mw loggingMiddleware) GetAccountDetails(ctx context.Context, req model.UserRequestDB)(res model.UserResponseDB, err error){

	defer func(begin time.Time){
		if err != nil{
			zap.Int64("timeTaken:", int64(time.Since(begin)))
			mw.logger.Error("Encounterd error processing request")
			mw.logger.Error(err.Error())
		}
		mw.logger.Info("able to pass through logging GetAccountDetails")
		mw.logger.Info("get details for email:"+ req.Email)
	}(time.Now())

	return mw.next.GetAccountDetails(ctx, req)
}

func (mw loggingMiddleware) SaveAccountDetails(ctx context.Context, req model.UserResponseDB)(res model.SaveResponseDB, err error){

	defer func(begin time.Time){
		if err != nil{
			zap.Int64("timeTaken:", int64(time.Since(begin)))
			mw.logger.Error("Encounterd error processing request")
			mw.logger.Error(err.Error())
		}
		mw.logger.Info("able to pass through logging SaveAccountDetails")
		mw.logger.Info("get details for email:"+ req.Email)
	}(time.Now())

	return mw.next.SaveAccountDetails(ctx, req)
}
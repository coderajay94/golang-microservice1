package base

import (
	"context"
	"errors"

	"github.com/coderajay94/microservice1/model"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AccountDetails endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		AccountDetails: MakeAccountDetails(s),
	}
}

func MakeAccountDetails(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.UserRequestDB)
		if !ok {
			return nil, errors.New("Invalid request")
		}
		return s.GetAccountDetails(ctx, req)
	}
}
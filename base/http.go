package base

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coderajay94/microservice1/model"
	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(e Endpoints, basepath string) http.Handler {

	r := chi.NewRouter()

	//r.Mount("/health", getHealth)

	//setup services routes
	r.Group(func(r chi.Router) {
		r.Route(basepath, func(r chi.Router) {
			r.Method(http.MethodPost, "/accountDetails", httptransport.NewServer(
				e.AccountDetails,
				decodeRequestAccountDetails,
				endcodeResponsAaccountDetails,
			))
		})
	})

	return r
}

func endcodeResponsAaccountDetails(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func decodeRequestAccountDetails(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req model.UserRequestDB
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func getHealth() string {
	return "200 OK"
}

package health

import (
	"net/http"

	"context"
	"encoding/json"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
)

// errorer describes the behavior of a request or response that can contain errors
type errorer interface {
	error() error
}

// MakeHandler builds a go-kit http transport and returns it
func MakeHandler(logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	e := makeHealthCheckEndpoint()

	healthHandler := kithttp.NewServer(
		e,
		decodeHealthCheckRequest,
		encodeHealthCheckResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/api/v1/health", healthHandler).Methods("GET")
	return r
}

// decodeHealthCheckRequest returns an empty healthCheck request because there are no params for this request
func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return healthCheckRequest{}, nil
}

// encodeHealthCheckResponse encodes any errors received from handling the request and returns
func encodeHealthCheckResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeError writes error headers if an error was received from a health check
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

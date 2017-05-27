package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// healthCheckRequest has no parameters, but we still generate an empty struct to represent it
type healthCheckRequest struct{}

// healthCheckResponse represents an HTTP response from the health endpoint containing any errors
type healthCheckResponse struct {
	Error error `json:"error,omitempty"`
}

// error is an implementation of the errorer interface allowing us to encode errors received from the service
func (r healthCheckResponse) error() error { return r.Error }

// makeHealthEndpoint returns a go-kit endpoint, wrapping the health response
func makeHealthCheckEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return healthCheckResponse{}, nil
	}
}

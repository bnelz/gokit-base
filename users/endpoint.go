package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// userCreateRequest represents an HTTP request from the client for user creation
type userCreateRequest struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	FavoriteColor string `json:"fav_color,omitempty"`
}

// userCreateResponse represents an HTTP response from our server for user creation
type userCreateResponse struct {
	ID    int   `json:"id,omitempty"`
	Error error `json:"error,omitempty"`
}

// error is the userCreateResponse errorer implementation
func (r userCreateResponse) error() error { return r.Error }

// makeCreateUserEndpoint generates a service endpoint for users
func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*userCreateRequest)

		id, err := s.CreateUser(req.ID, req.FirstName, req.LastName, req.FavoriteColor)

		return userCreateResponse{ID: id, Error: err}, nil
	}
}

// userReadRequest represents an HTTP request to read a single user from the client
type userReadRequest struct {
	ID int `json:"id"`
}

// userReadResponse represents an HTTP response containing a user or the error when fetching
type userReadResponse struct {
	User  User  `json:"user,omitempty"`
	Error error `json:"error,omitempty"`
}

// error is the userReadResponse errorer implementation
func (r userReadResponse) error() error { return r.Error }

func makeReadUserEindpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userReadRequest)
		u, err := s.ReadUser(req.ID)
		return userReadResponse{User: u, Error: err}, nil
	}
}

// userUpdateColorRequest represents an HTTP request from the client to update a user's favorite color
type userUpdateColorRequest struct {
	ID            int    `json:"id"`
	FavoriteColor string `json:"favorite_color"`
}

// userUpdateColorResponse represents an HTTP response from the server notifying the client of the update status
type userUpdateColorResponse struct {
	Error error `json:"error,omitempty"`
}

// error is an errorer implementation for userUpdateColorResponse
func (r userUpdateColorResponse) error() error { return r.Error }

func makeUpdateUserColorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userUpdateColorRequest)
		err := s.UpdateUserColor(req.ID, req.FavoriteColor)
		return userUpdateColorResponse{Error: err}, nil
	}
}

// userReadAllRequest represents an HTTP request from the client to get all users
type userReadAllRequest struct{}

// userReadAllResponse represents an HTTP response from the server listing all users
type userReadAllResponse struct {
	Users []*User `json:"users,omitempty"`
	Error error   `json:"error,omitempty"`
}

// error is an errorer implementation for userReadAllResponse
func (r userReadAllResponse) error() error { return r.Error }

// makeReadAllUsersEndpoint creates an HTTP endpoint for retrieving all users
func makeReadAllUsersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users := s.Users()
		return userReadAllResponse{Users: users, Error: nil}, nil
	}
}

package users

import (
	"net/http"

	"context"
	"encoding/json"

	"io/ioutil"

	"strconv"

	errs "github.com/bnelz/gokit-base/errors"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

func MakeHandler(us Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	// Define all endpoints
	create := makeCreateUserEndpoint(us)
	read := makeReadUserEindpoint(us)
	update := makeUpdateUserColorEndpoint(us)
	list := makeReadAllUsersEndpoint(us)

	createHandler := kithttp.NewServer(
		create,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		opts...,
	)

	readHandler := kithttp.NewServer(
		read,
		decodeReadUserRequest,
		encodeReadUserResponse,
		opts...,
	)

	updateHandler := kithttp.NewServer(
		update,
		decodeUpdateUserRequest,
		encodeUpdateUserResponse,
		opts...,
	)

	listHandler := kithttp.NewServer(
		list,
		decodeListUsersRequest,
		encodeListUsersResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/api/v1/users", listHandler).Methods("GET")
	r.Handle("/api/v1/users", createHandler).Methods("POST")
	r.Handle("/api/v1/users/{id}", readHandler).Methods("GET")
	r.Handle("/api/v1/users/{id}", updateHandler).Methods("PUT")

	return r
}

func decodeRequest(to interface{}, r *http.Request) (interface{}, error) {
	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &to)
	return to, err
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	d, err := json.Marshal(response)
	if err != nil {
		encodeError(ctx, err, w)
	}
	_, err = w.Write(d)
	return err
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userCreateRequest
	return decodeRequest(&req, r)
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	res := response.(userCreateResponse)
	return encodeResponse(ctx, w, res)
}

func decodeReadUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	req := userReadRequest{
		ID: id,
	}
	return req, err
}

func encodeReadUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	res := response.(userReadResponse)
	return encodeResponse(ctx, w, res)
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userUpdateColorRequest
	return decodeRequest(req, r)
}

func encodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	res := response.(userUpdateColorResponse)
	return encodeResponse(ctx, w, res)
}

func decodeListUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := userReadAllRequest{}
	return req, nil
}

func encodeListUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	res := response.(userReadAllResponse)
	return encodeResponse(ctx, w, res)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case errs.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case errs.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

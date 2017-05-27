package errors

import "errors"

var (
	ErrInvalidArgument = errors.New("Invalid function argument(s)")
	ErrUserNotFound    = errors.New("User not found")
)

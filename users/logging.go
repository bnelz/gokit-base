package users

import (
	"time"

	"github.com/go-kit/kit/log"
)

// encapsulates logging for our service
type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService generates a new logging service
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

// CreateUser wraps the user service method with logging metadata we want to capture and defers the call
func (s *loggingService) CreateUser(id int, fname string, lname string, color string) (retID int, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"context_method", "CreateUser",
			"context_id", id,
			"context_fname", fname,
			"context_lname", lname,
			"context_color", color,
			"context_elapsed_time", time.Since(begin),
			"message", err,
		)
	}(time.Now())
	return s.Service.CreateUser(id, fname, lname, color)
}

package users

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

// instrumentingService encapsulates metric aggregation for our files service
type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService generates a new instance of our instrumented files service
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) CreateUser(id int, fname string, lname string, color string) (int, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateUser").Add(1)
		s.requestLatency.With("method", "CreateUser").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.CreateUser(id, fname, lname, color)
}

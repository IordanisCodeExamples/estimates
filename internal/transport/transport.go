package transport

import (
	"context"
)

type Service interface {
	GetEstimates(ctx context.Context, limit, lengthOfInterview int32) (bool, string, error)
}

// Logger defines the interface for logging functionalities required within the service.
type Logger interface {
	Info(ctx context.Context, msg string, fields ...map[string]interface{})
	Error(ctx context.Context, msg string, fields ...map[string]interface{})
}

package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type BadGatewayError struct {
	APIError
}

// NewBadGateway creates a new BadGatewayError with corresponding status codes:
//
// - HTTP: 502
//
// - GRPC: 14 .
func NewBadGateway(
	msg string,
	enum string,
) *BadGatewayError {
	return &BadGatewayError{
		APIError: APIError{
			msg:  msg,
			grpc: codes.Unavailable,
			http: http.StatusBadGateway,
			enum: enum,
		},
	}
}

func IsBadGateway(err error) bool {
	_, ok := AsBadGateway(err)
	return ok
}

func AsBadGateway(err error) (*BadGatewayError, bool) {
	var t *BadGatewayError
	return t, errors.As(err, &t)
}

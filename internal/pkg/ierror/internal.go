package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type InternalError struct {
	APIError
}

// NewInternal creates a new InternalError with corresponding status codes:
//
// - HTTP: 500
//
// - GRPC: 13 .
func NewInternal(msg, enum string) *BadGatewayError {
	return &BadGatewayError{
		APIError: APIError{
			msg:  msg,
			grpc: codes.Internal,
			http: http.StatusInternalServerError,
			enum: enum,
		},
	}
}

func IsInternalError(err error) bool {
	_, ok := AsInternalError(err)
	return ok
}

func AsInternalError(err error) (*InternalError, bool) {
	var t *InternalError
	return t, errors.As(err, &t)
}

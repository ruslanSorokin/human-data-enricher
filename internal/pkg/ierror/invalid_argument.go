package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type InvalidArgumentError struct {
	APIError
}

var _ APIErrorI = (*InvalidArgumentError)(nil)

// NewInvalidArgument creates a new InvalidArgumentError with corresponding status codes:
//
// - HTTP: 404
//
// - GRPC: 5 .
func NewInvalidArgument(msg, enum string) *InvalidArgumentError {
	return &InvalidArgumentError{
		APIError: APIError{
			msg:  msg,
			grpc: codes.InvalidArgument,
			http: http.StatusBadRequest,
			enum: enum,
		},
	}
}

func IsInvalidArgument(err error) bool {
	_, ok := AsInvalidArgument(err)
	return ok
}

func AsInvalidArgument(err error) (*InvalidArgumentError, bool) {
	var t *InvalidArgumentError
	return t, errors.As(err, &t)
}

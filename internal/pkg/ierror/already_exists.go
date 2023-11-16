package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type AlreadyExistsErrorI interface {
	APIErrorI

	DuplicateID() (string, bool)
}

type AlreadyExistsError struct {
	duplicateID string

	APIError
}

var _ AlreadyExistsErrorI = (*AlreadyExistsError)(nil)

// NewInvalidArgument creates a new InvalidArgumentError with corresponding status codes:
//
// - HTTP: 409
//
// - GRPC: 6 .
func NewAlreadyExists(
	msg string,
	enum string,
) *AlreadyExistsError {
	return &AlreadyExistsError{
		duplicateID: "",
		APIError: APIError{
			msg:  msg,
			grpc: codes.AlreadyExists,
			http: http.StatusConflict,
			enum: enum,
		},
	}
}

// DuplicateID returns duplicate ID if field is populated.
func (e AlreadyExistsError) DuplicateID() (string, bool) {
	return e.duplicateID, e.duplicateID != ""
}

type InstantiatedAlreadyExistsError struct {
	AlreadyExistsError

	parent *AlreadyExistsError
}

func (e *AlreadyExistsError) Instantiate() *InstantiatedAlreadyExistsError {
	return &InstantiatedAlreadyExistsError{
		AlreadyExistsError: *NewAlreadyExists(e.msg, e.enum),
		parent:             e,
	}
}

// WithDuplicateID returns a copy of AlreadyExistsError with populated
// "duplicateID" field.
func (e *InstantiatedAlreadyExistsError) WithDuplicateID(
	dupID string,
) *InstantiatedAlreadyExistsError {
	e.duplicateID = dupID
	return e
}

func (e InstantiatedAlreadyExistsError) Is(target error) bool {
	var t *AlreadyExistsError
	return errors.As(target, &t) && t.APIError == e.parent.APIError
}

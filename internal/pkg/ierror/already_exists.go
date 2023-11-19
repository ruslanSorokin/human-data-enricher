package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type AlreadyExistsErrorI interface {
	APIErrorI

	DuplicateID() (string, bool)
	UniqueProperty() (string, bool)
}

type AlreadyExistsError struct {
	duplicateID    string
	uniqueProperty string

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
		duplicateID:    "",
		uniqueProperty: "",
		APIError: APIError{
			msg:  msg,
			grpc: codes.AlreadyExists,
			http: http.StatusConflict,
			enum: enum,
		},
	}
}

// DuplicateID returns duplicate ID if the field is populated.
func (e AlreadyExistsError) DuplicateID() (string, bool) {
	return e.duplicateID, e.duplicateID != ""
}

// UniqueProperty returns uniqueness property if the field is populated.
func (e AlreadyExistsError) UniqueProperty() (string, bool) {
	return e.uniqueProperty, e.uniqueProperty != ""
}

type InstantiatedAlreadyExistsError struct {
	parent *AlreadyExistsError

	AlreadyExistsError
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

// WithUniqueProperty returns a copy of AlreadyExistsError with populated
// "uniqueProperty" field.
func (e *InstantiatedAlreadyExistsError) WithUniqueProperty(
	prop string,
) *InstantiatedAlreadyExistsError {
	e.uniqueProperty = prop
	return e
}

func (e InstantiatedAlreadyExistsError) Is(target error) bool {
	var t *AlreadyExistsError
	return errors.As(target, &t) && t.APIError == e.parent.APIError
}

func IsAlreadyExists(err error) bool {
	_, ok := AsAlreadyExists(err)
	return ok
}

func AsAlreadyExists(err error) (*AlreadyExistsError, bool) {
	var t *AlreadyExistsError
	return t, errors.As(err, &t)
}

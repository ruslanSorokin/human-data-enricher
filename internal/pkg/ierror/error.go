// Package provides functionality to create API-scope errors for CRUD-like
// applications.
//
// term 'static' means compile-time or immutable error such as:
//
//	var ErrAlreadyExists = ierror.NewAlreadyExists("user with this login already exists", "USER_ALREADY_EXISTS")
//
// whearas 'dynamic' means that you want to create a copy of a 'static' error
// and populate it with some information:
//
//	if (...){
//		 return nil, ierror.InstantiateAlreadyExists(ErrAlreadyExists, id)
//	}
//
// All 'NewXXX' ctors return 'static' errors. And all 'static' errors that can
// be populated with some request-specific parameters and transformed into
// 'dynamic' ones, can be transformed into them by 'InstantiateXXX' method.
package ierror

import (
	"errors"

	"google.golang.org/grpc/codes"
)

type GRPCConvertible interface {
	// ToGRPC returns GRPC status code.
	ToGRPC() codes.Code
}

type HTTPConvertible interface {
	// ToHTTP returns HTTP status code.
	ToHTTP() int
}

type EnumConvertible interface {
	// ToEnum returns enum.
	ToEnum() string
}

type APIErrorI interface {
	error

	GRPCConvertible
	HTTPConvertible
	EnumConvertible
}

type PropertyError struct {
	property string

	InvalidArgumentError
}

type APIError struct {
	msg  string
	enum string
	grpc codes.Code
	http int
}

var _ APIErrorI = (*APIError)(nil)

// New creates new a APIError.
func New(msg string, grpc codes.Code, http int, enum string) *APIError {
	return &APIError{msg: msg, grpc: grpc, http: http, enum: enum}
}

func (e *APIError) Error() string { return e.msg }

func (e *APIError) ToGRPC() codes.Code { return e.grpc }

func (e *APIError) ToHTTP() int { return e.http }

func (e *APIError) ToEnum() string { return e.enum }

type InstantiatedAPIError struct {
	parent *APIError

	APIError
}

func (e *APIError) Instantiate() *InstantiatedAPIError {
	return &InstantiatedAPIError{
		APIError: *New(e.msg, e.grpc, e.http, e.enum),
		parent:   e,
	}
}

func (e *InstantiatedAPIError) WithEnum(
	enum string,
) *InstantiatedAPIError {
	e.enum = enum
	return e
}

func (e *InstantiatedAPIError) WithGRPCStCode(
	code codes.Code,
) *InstantiatedAPIError {
	e.grpc = code
	return e
}

func (e *InstantiatedAPIError) WithHTTPStCode(
	code int,
) *InstantiatedAPIError {
	e.http = code
	return e
}

func (e InstantiatedAPIError) Is(target error) bool {
	var t *APIError
	return errors.As(target, &t) && t == e.parent
}

func IsAPIError(err error) bool {
	_, ok := AsAPIError(err)
	return ok
}

func AsAPIError(err error) (*APIError, bool) {
	var t *APIError
	return t, errors.As(err, &t)
}

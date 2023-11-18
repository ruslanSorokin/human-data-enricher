package ierror

import (
	"errors"
	"fmt"
)

type InvalidPropertyErrorI interface {
	APIErrorI

	Violation() (string, bool)
}

type InvalidPropertyError struct {
	violation string

	PropertyError
}

var _ InvalidPropertyErrorI = (*InvalidPropertyError)(nil)

// NewInvalidProperty creates a new InvalidPropertyError with corresponding status codes:
//
// - HTTP: 409
//
// - GRPC: 6 .
func NewInvalidProperty(prop, enum string) *InvalidPropertyError {
	msg := fmt.Sprintf("invalid property: %s", prop)
	return &InvalidPropertyError{
		violation: "",
		PropertyError: PropertyError{
			InvalidArgumentError: *NewInvalidArgument(msg, enum),
			property:             prop,
		},
	}
}

// Violation returns property's violation.
func (e InvalidPropertyError) Violation() (string, bool) {
	return e.violation, e.violation != ""
}

type InstantiatedInvalidPropertyError struct {
	parent *InvalidPropertyError

	InvalidPropertyError
}

func (e *InvalidPropertyError) Instantiate() *InstantiatedInvalidPropertyError {
	return &InstantiatedInvalidPropertyError{
		InvalidPropertyError: *NewInvalidProperty(e.property, e.enum),
		parent:               e,
	}
}

// WithViolation creates a new InvalidPropertyError from "err" with
// populated "violation" field.
func (e *InstantiatedInvalidPropertyError) WithViolation(
	vio string,
) *InvalidPropertyError {
	return &InvalidPropertyError{
		PropertyError: e.PropertyError,
		violation:     vio,
	}
}

func (e InstantiatedInvalidPropertyError) Is(target error) bool {
	var t *InvalidPropertyError
	return errors.As(target, &t) && t == e.parent
}

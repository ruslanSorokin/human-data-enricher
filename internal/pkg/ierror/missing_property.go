package ierror

import (
	"errors"
	"fmt"
)

// MissingPropertyError means that some mondatory property is missing.
type MissingPropertyError struct {
	PropertyError
}

var _ APIErrorI = (*MissingPropertyError)(nil)

// NewMissingProperty creates a new MissingPropertyError with corresponding status codes:
//
// - HTTP: 409
//
// - GRPC: 6 .
func NewMissingProperty(prop, enum string) error {
	msg := fmt.Sprintf("missing property: %s", prop)
	return &MissingPropertyError{
		PropertyError: PropertyError{
			InvalidArgumentError: *NewInvalidArgument(msg, enum),
			property:             prop,
		},
	}
}

func IsMissingProperty(err error) bool {
	_, ok := AsMissingProperty(err)
	return ok
}

func AsMissingProperty(err error) (*MissingPropertyError, bool) {
	var t *MissingPropertyError
	return t, errors.As(err, &t)
}

package ierror

import "fmt"

type PropertyError struct {
	property string
}

type MissingPropertyError PropertyError

var _ error = (*MissingPropertyError)(nil)

// NewMissingProperty returns MissingPropertyError.
func NewMissingProperty(prop string) error {
	return &MissingPropertyError{
		property: prop,
	}
}

func (e MissingPropertyError) Error() string {
	return fmt.Sprintf("missing %s", e.property)
}

type InvalidPropertyError struct {
	PropertyError
}

// NewInvalidProperty returns InvalidPropertyError.
func NewInvalidProperty(prop string) error {
	return &InvalidPropertyError{
		PropertyError: PropertyError{
			property: prop,
		},
	}
}

var _ error = (*InvalidPropertyError)(nil)

func (e InvalidPropertyError) Error() string {
	return fmt.Sprintf("invalid %s", e.property)
}

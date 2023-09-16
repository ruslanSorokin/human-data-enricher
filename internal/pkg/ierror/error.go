package ierror

import "fmt"

type PropertyError struct {
	propetry string
}

type MissingPropertyError PropertyError

var _ error = (*MissingPropertyError)(nil)

// NewMissingProperty returns MissingPropertyError.
func NewMissingProperty(prop string) error {
	return &MissingPropertyError{
		propetry: prop,
	}
}

func (e MissingPropertyError) Error() string {
	return fmt.Sprintf("missing %s", e.propetry)
}

type InvalidPropertyError struct {
	PropertyError
}

// NewInvalidProperty returns InvalidPropertyError.
func NewInvalidProperty(prop string) error {
	return &InvalidPropertyError{
		PropertyError: PropertyError{
			propetry: prop,
		},
	}
}

var _ error = (*InvalidPropertyError)(nil)

func (e InvalidPropertyError) Error() string {
	return fmt.Sprintf("invalid %s", e.propetry)
}

package ierror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

func TestInstantiateInvalidProperty(t *testing.T) {
	assert := assert.New(t)

	nameStatic := ierror.NewInvalidProperty(
		"invalid name",
		"INVALID_NAME",
	)
	nameDynamic := nameStatic.Instantiate().
		WithViolation("contains numeric character")

	assert.False(errors.Is(nameStatic, nameDynamic))

	surnameStatic := ierror.NewInvalidProperty(
		"invalid surname",
		"INVALID_SURNAME",
	)
	surnameDynamic := surnameStatic.Instantiate().
		WithViolation("contains numeric character")

	assert.False(errors.Is(surnameStatic, surnameDynamic))

	assert.False(errors.Is(surnameStatic, nameDynamic))
	assert.False(errors.Is(nameStatic, surnameDynamic))

	assert.False(errors.Is(nameStatic, errors.New("random error here")))
}

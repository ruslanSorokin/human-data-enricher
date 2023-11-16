package ivalidator_test

import (
	"log/slog"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation/ivalidator"
	validationtest "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation/test"
)

type Suite struct {
	*suite.Suite
	*validationtest.VSuite
}

func TestValidator(t *testing.T) {
	s := &suite.Suite{Assertions: assert.New(t)}

	vtor := ivalidator.NewValidator(
		slog.Default(),
		validator.New(validator.WithRequiredStructEnabled()))

	vs := validationtest.NewSuite(s, vtor)

	suite.Run(t, &Suite{
		VSuite: vs,
		Suite:  s,
	})
}

package validationtest

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation"
)

type VSuite struct {
	suite.TestingSuite

	vtor validation.ValidatorI
}

func NewSuite(s suite.TestingSuite, v validation.ValidatorI) *VSuite {
	return &VSuite{
		TestingSuite: s,
		vtor:         v,
	}
}

package test_nationality

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

type ProviderSuite struct {
	suite.TestingSuite

	Provider provider.NationalityProviderI
}

func NewSuite(
	s suite.TestingSuite,
	p provider.NationalityProviderI,
) *ProviderSuite {
	return &ProviderSuite{TestingSuite: s, Provider: p}
}

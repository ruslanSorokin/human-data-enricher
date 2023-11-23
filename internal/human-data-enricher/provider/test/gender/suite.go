package test_gender

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

type ProviderSuite struct {
	suite.TestingSuite

	Provider provider.GenderProviderI
}

func NewSuite(s suite.TestingSuite, p provider.GenderProviderI) *ProviderSuite {
	return &ProviderSuite{TestingSuite: s, Provider: p}
}

package test_age

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

type ProviderSuite struct {
	suite.TestingSuite

	Provider provider.AgeProviderI
}

func NewSuite(s suite.TestingSuite, p provider.AgeProviderI) *ProviderSuite {
	return &ProviderSuite{TestingSuite: s, Provider: p}
}

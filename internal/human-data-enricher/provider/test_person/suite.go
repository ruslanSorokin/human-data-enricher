package test_person

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

type PSuite struct {
	suite.TestingSuite

	Provider provider.PersonProviderI
}

func NewSuite(s suite.TestingSuite, p provider.PersonProviderI) *PSuite {
	return &PSuite{TestingSuite: s, Provider: p}
}

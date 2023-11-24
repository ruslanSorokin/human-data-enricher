package ifiber_test

import (
	"flag"
	"log/slog"
	"testing"

	"github.com/gofiber/fiber/v2"
	ifiber "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/age"
	test_age "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test/age"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_age.ProviderSuite

	client *fiber.Client
}

func NewIntegrationSuite(s *suite.Suite) *IntegrationSuite {
	return &IntegrationSuite{
		Suite:         s,
		ProviderSuite: nil,
		client:        nil,
	}
}

func TestIntegration_Fiber_AgeGateway(t *testing.T) {
	flag.Parse()
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, NewIntegrationSuite(&suite.Suite{Assertions: assert.New(t)}))
}

func (s *IntegrationSuite) SetupSuite() {
	cl := fiber.AcquireClient()
	s.client = cl

	log := slog.Default()

	gw := ifiber.NewAgeGateway(log, cl)
	s.ProviderSuite = test_age.NewSuite(s, gw)
}

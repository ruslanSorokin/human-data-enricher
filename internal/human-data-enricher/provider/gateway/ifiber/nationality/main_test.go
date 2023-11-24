package ifiber_test

import (
	"flag"
	"log/slog"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ifiber "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/nationality"
	test_nationality "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test/nationality"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_nationality.ProviderSuite

	client *fiber.Client
}

func NewIntegrationSuite(s *suite.Suite) *IntegrationSuite {
	return &IntegrationSuite{
		Suite:         s,
		ProviderSuite: nil,
		client:        nil,
	}
}

func TestIntegration_Fiber_NationalityGateway(t *testing.T) {
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

	gw := ifiber.NewNationalityGateway(log, cl)
	s.ProviderSuite = test_nationality.NewSuite(s, gw)
}

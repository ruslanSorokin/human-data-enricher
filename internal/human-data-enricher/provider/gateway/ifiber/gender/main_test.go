package ifiber_test

import (
	"flag"
	"log/slog"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ifiber "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/gender"
	test_gender "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test/gender"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_gender.ProviderSuite

	client *fiber.Client
}

func NewIntegrationSuite(s *suite.Suite) *IntegrationSuite {
	return &IntegrationSuite{
		Suite:         s,
		ProviderSuite: nil,
		client:        nil,
	}
}

func TestIntegration_Fiber_GenderGateway(t *testing.T) {
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

	gw := ifiber.NewGenderGateway(log, cl)
	s.ProviderSuite = test_gender.NewSuite(s, gw)
}

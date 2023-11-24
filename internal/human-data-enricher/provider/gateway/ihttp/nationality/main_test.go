package ihttp_test

import (
	"flag"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ihttp "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ihttp/nationality"
	test_nationality "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test/nationality"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_nationality.ProviderSuite

	client *http.Client
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
	cl := &http.Client{}
	s.client = cl

	log := slog.Default()

	gw := ihttp.NewNationalityGateway(log, cl)
	s.ProviderSuite = test_nationality.NewSuite(s, gw)
}

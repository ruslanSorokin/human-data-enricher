package ihttp_test

import (
	"flag"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ihttp "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ihttp/gender"
	test_gender "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test/gender"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_gender.ProviderSuite

	client *http.Client
}

func NewIntegrationSuite(s *suite.Suite) *IntegrationSuite {
	return &IntegrationSuite{
		Suite:         s,
		ProviderSuite: nil,
		client:        nil,
	}
}

func TestIntegration_HTTP_GenderGateway(t *testing.T) {
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

	gw := ihttp.NewGenderGateway(log, cl)
	s.ProviderSuite = test_gender.NewSuite(s, gw)
}

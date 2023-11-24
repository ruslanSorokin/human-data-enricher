package ihttp

import (
	"log/slog"
	"net/http"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	ihttp_age "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ihttp/age"
	ihttp_gender "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ihttp/gender"
	ihttp_nationality "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ihttp/nationality"
)

var ErrBadAgeGateway = provider.ErrBadAgeGateway

type AgeGateway ihttp_age.AgeGateway

func NewAgeGateway(l *slog.Logger, c *http.Client) *ihttp_age.AgeGateway {
	return ihttp_age.NewAgeGateway(l, c)
}

var ErrBadGenderGateway = provider.ErrBadGenderGateway

type GenderGateway ihttp_gender.GenderGateway

func NewGenderGateway(
	l *slog.Logger,
	c *http.Client,
) *ihttp_gender.GenderGateway {
	return ihttp_gender.NewGenderGateway(l, c)
}

var ErrBadNationalityGateway = provider.ErrBadNationalityGateway

type NationalityGateway ihttp_nationality.NationalityGateway

func NewNationalityGateway(
	l *slog.Logger,
	c *http.Client,
) *ihttp_nationality.NationalityGateway {
	return ihttp_nationality.NewNationalityGateway(l, c)
}

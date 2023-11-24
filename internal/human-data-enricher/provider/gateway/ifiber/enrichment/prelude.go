package ifiber

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	ifiber_age "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/age"
	ifiber_gender "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/gender"
	ifiber_nationality "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/gateway/ifiber/nationality"
)

var ErrBadAgeGateway = provider.ErrBadAgeGateway

type AgeGateway ifiber_age.AgeGateway

func NewAgeGateway(
	l *slog.Logger,
	c *fiber.Client,
) *ifiber_age.AgeGateway {
	return ifiber_age.NewAgeGateway(l, c)
}

var ErrBadGenderGateway = provider.ErrBadGenderGateway

type GenderGateway ifiber_gender.GenderGateway

func NewGenderGateway(
	l *slog.Logger,
	c *fiber.Client,
) *ifiber_gender.GenderGateway {
	return ifiber_gender.NewGenderGateway(l, c)
}

var ErrBadNationalityGateway = provider.ErrBadNationalityGateway

type NationalityGateway ifiber_nationality.NationalityGateway

func NewNationalityGateway(
	l *slog.Logger,
	c *fiber.Client,
) *ifiber_nationality.NationalityGateway {
	return ifiber_nationality.NewNationalityGateway(l, c)
}

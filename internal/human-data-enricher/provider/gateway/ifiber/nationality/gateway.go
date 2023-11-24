package ifiber

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

const (
	baseURL    = "https://api.nationalize.io"
	qParamName = "name"
	method     = "GET"
)

type NationalityGateway struct {
	log    *slog.Logger
	client *fiber.Client
}

func NewNationalityGateway(
	l *slog.Logger,
	c *fiber.Client,
) *NationalityGateway {
	return &NationalityGateway{log: l, client: c}
}

var _ provider.NationalityProviderI = (*NationalityGateway)(nil)

func (g *NationalityGateway) NationalityByName(
	ctx context.Context,
	n string,
) (string, error) {
	resp := fiber.AcquireResponse()
	defer fiber.ReleaseResponse(resp)

	args := fiber.AcquireArgs()
	defer fiber.ReleaseArgs(args)
	args.Set(qParamName, n)

	a := g.client.
		Get(baseURL).
		SetResponse(resp).
		QueryString(args.String()).
		ContentType("application/json")
	defer fiber.ReleaseAgent(a)

	if err := a.Parse(); err != nil {
		g.log.Error("bad attempt to create a request",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadNationalityGateway, err)
	}

	stCode, body, errs := a.String()
	if errs != nil {
		g.log.Error("bad attempt to make a request",
			"body", body,
			"stCode", stCode,
			"error", errs)
		return "", fmt.Errorf(
			"%w: %w",
			provider.ErrBadNationalityGateway,
			errors.Join(errs...),
		)
	}

	var respBody response

	if err := json.NewDecoder(strings.NewReader(body)).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadNationalityGateway, err)
	}

	if len(respBody.Countries) > 0 {
		return respBody.Countries[0].CountryID, nil
	}

	return "", provider.ErrBadNationalityGateway
}

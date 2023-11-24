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
	baseURL    = "https://api.agify.io"
	qParamName = "name"
	method     = "GET"
)

type AgeGateway struct {
	log    *slog.Logger
	client *fiber.Client
}

func NewAgeGateway(l *slog.Logger, c *fiber.Client) *AgeGateway {
	return &AgeGateway{log: l, client: c}
}

var _ provider.AgeProviderI = (*AgeGateway)(nil)

func (g *AgeGateway) AgeByName(
	ctx context.Context,
	n string,
) (int, error) {
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
		return 0, fmt.Errorf("%w: %w", provider.ErrBadAgeGateway, err)
	}

	stCode, body, errs := a.String()
	if errs != nil {
		g.log.Error("bad attempt to make a request",
			"body", body,
			"stCode", stCode,
			"error", errs)
		return 0, fmt.Errorf(
			"%w: %w",
			provider.ErrBadAgeGateway,
			errors.Join(errs...),
		)
	}

	var respBody struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Count int    `json:"count"`
	}

	if err := json.NewDecoder(strings.NewReader(body)).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"error", err)
		return 0, fmt.Errorf("%w: %w", provider.ErrBadAgeGateway, err)
	}

	return respBody.Age, nil
}

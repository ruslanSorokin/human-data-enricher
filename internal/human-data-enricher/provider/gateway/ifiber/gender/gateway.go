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
	baseURL    = "https://api.genderize.io"
	qParamName = "name"
	method     = "GET"
)

type GenderGateway struct {
	log    *slog.Logger
	client *fiber.Client
}

func NewGenderGateway(l *slog.Logger, c *fiber.Client) *GenderGateway {
	return &GenderGateway{log: l, client: c}
}

var _ provider.GenderProviderI = (*GenderGateway)(nil)

func (g *GenderGateway) GenderByName(
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
		return "", fmt.Errorf("%w: %w", provider.ErrBadGenderGateway, err)
	}

	stCode, body, errs := a.String()
	if errs != nil {
		g.log.Error("bad attempt to make a request",
			"body", body,
			"stCode", stCode,
			"error", errs)
		return "", fmt.Errorf(
			"%w: %w",
			provider.ErrBadGenderGateway,
			errors.Join(errs...),
		)
	}

	var respBody struct {
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Count       int     `json:"count"`
		Probability float64 `json:"probability"`
	}

	if err := json.NewDecoder(strings.NewReader(body)).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadGenderGateway, err)
	}

	return respBody.Gender, nil
}

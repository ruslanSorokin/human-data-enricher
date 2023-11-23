package ihttp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

const (
	baseURL    = "https://api.agify.io"
	qParamName = "name"
	method     = "GET"
)

type AgeGateway struct {
	log    *slog.Logger
	client *http.Client
}

func NewAgeGateway(l *slog.Logger, c *http.Client) *AgeGateway {
	return &AgeGateway{log: l, client: c}
}

var _ provider.AgeProviderI = (*AgeGateway)(nil)

func (g *AgeGateway) AgeByName(ctx context.Context, n string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, baseURL, nil)
	if err != nil {
		g.log.Error("bad attempt to create new request",
			"error", err)
		return 0, fmt.Errorf("%w: %w", provider.ErrBadAgeGateway, err)
	}

	q := req.URL.Query()
	q.Add(qParamName, n)
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		g.log.Error("bad attempt to execute request",
			"name", n,
			"error", err)
		return 0, fmt.Errorf("%w: %w", provider.ErrBadAgeGateway, err)
	}

	var respBody struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Count int    `json:"count"`
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"respBody", respBody,
			"error", err)
		return 0, fmt.Errorf("%w: %w", provider.ErrBadAgeGateway, err)
	}

	return respBody.Age, nil
}

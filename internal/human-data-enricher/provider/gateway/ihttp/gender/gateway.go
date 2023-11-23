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
	baseURL    = "https://api.genderize.io"
	qParamName = "name"
	method     = "GET"
)

type GenderGateway struct {
	log    *slog.Logger
	client *http.Client
}

func NewGenderGateway(l *slog.Logger, c *http.Client) *GenderGateway {
	return &GenderGateway{log: l, client: c}
}

var _ provider.GenderProviderI = (*GenderGateway)(nil)

func (g *GenderGateway) GenderByName(ctx context.Context, n string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, baseURL, nil)
	if err != nil {
		g.log.Error("bad attempt to create new request",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadGenderGateway, err)
	}

	q := req.URL.Query()
	q.Add(qParamName, n)
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		g.log.Error("bad attempt to execute request",
			"name", n,
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadGenderGateway, err)
	}

	var respBody struct {
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Count       int     `json:"count"`
		Probability float64 `json:"probability"`
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"respBody", respBody,
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadGenderGateway, err)
	}

	return respBody.Gender, nil
}

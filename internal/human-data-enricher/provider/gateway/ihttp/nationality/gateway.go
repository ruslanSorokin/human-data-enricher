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
	baseURL    = "https://api.nationalize.io"
	qParamName = "name"
	method     = "GET"
)

type NationalityGateway struct {
	log    *slog.Logger
	client *http.Client
}

func NewNationalityGateway(l *slog.Logger, c *http.Client) *NationalityGateway {
	return &NationalityGateway{log: l, client: c}
}

var _ provider.NationalityProviderI = (*NationalityGateway)(nil)

func (g *NationalityGateway) NationalityByName(
	ctx context.Context,
	n string,
) (string, error) {
	req, err := http.NewRequestWithContext(ctx, method, baseURL, nil)
	if err != nil {
		g.log.Error("bad attempt to create request",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadNationalityGateway, err)
	}

	q := req.URL.Query()
	q.Add(qParamName, n)
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		g.log.Error("bad attempt to make a request",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadNationalityGateway, err)
	}

	var respBody response

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).DecodeContext(ctx, &respBody); err != nil {
		g.log.Error("bad attempt to decode response body",
			"error", err)
		return "", fmt.Errorf("%w: %w", provider.ErrBadNationalityGateway, err)
	}

	if len(respBody.Countries) > 0 {
		return respBody.Countries[0].CountryID, nil
	}

	return "", provider.ErrBadNationalityGateway
}

package provider

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

var ErrBadNationalityGateway = ierror.NewBadGateway(
	"bad nationality gateway",
	"BAD_GATEWAY",
)

type NationalityProviderI interface {
	NationalityByName(context.Context, string) (string, error)
}

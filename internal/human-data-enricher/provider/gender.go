package provider

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

var ErrBadGenderGateway = ierror.NewBadGateway(
	"bad gender gateway",
	"BAD_GATEWAY",
)

type GenderProviderI interface {
	GenderByName(context.Context, string) (string, error)
}

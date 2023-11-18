package provider

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

var ErrBadAgeGateway = ierror.NewBadGateway(
	"bad age gateway",
	"BAD_GATEWAY",
)

type AgeProviderI interface {
	AgeByName(context.Context, string) (int, error)
}

package person_service

import (
	"log/slog"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation"
)

type ServiceI interface {
	CreatorI
	DeleterI
	UpdaterI
	GetterI

	ListGetterI
}

type PersonService struct {
	log     *slog.Logger
	vtor    validation.ValidatorI
	storage provider.PersonProviderI
}

var _ ServiceI = (*PersonService)(nil)

func New(
	log *slog.Logger,
	validator validation.ValidatorI,
	storage provider.PersonProviderI,
) *PersonService {
	return &PersonService{
		log:     log,
		vtor:    validator,
		storage: storage,
	}
}

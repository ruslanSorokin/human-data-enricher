package person_enricher

import (
	"log/slog"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	person_service "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person"
)

type EnricherI interface {
	ByNameEnricherI
}

type Enricher struct {
	log           *slog.Logger
	personCreator person_service.CreatorI

	ageProvider         provider.AgeProviderI
	genderProvider      provider.GenderProviderI
	nationalityProvider provider.NationalityProviderI
}

var _ EnricherI = (*Enricher)(nil)

func New(
	log *slog.Logger,
	personCreator person_service.CreatorI,
	ageProvider provider.AgeProviderI,
	genderProvider provider.GenderProviderI,
	nationalityProvider provider.NationalityProviderI,
) *Enricher {
	return &Enricher{
		log:                 log,
		personCreator:       personCreator,
		ageProvider:         ageProvider,
		genderProvider:      genderProvider,
		nationalityProvider: nationalityProvider,
	}
}

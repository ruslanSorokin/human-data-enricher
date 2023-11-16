package service

import (
	"log/slog"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	person_enricher "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/enricher"
	person_service "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation"
)

var (
	ErrMissingPersonName        = person_service.ErrMissingName
	ErrMissingPersonSurname     = person_service.ErrMissingSurname
	ErrMissingPersonAge         = person_service.ErrMissingAge
	ErrMissingPersonGender      = person_service.ErrMissingGender
	ErrMissingPersonNationality = person_service.ErrMissingNationality
	ErrInvalidPersonName        = person_service.ErrInvalidName
	ErrInvalidPersonSurname     = person_service.ErrInvalidSurname
	ErrInvalidPersonMiddleName  = person_service.ErrInvalidMiddleName
	ErrInvalidPersonAge         = person_service.ErrInvalidAge
	ErrInvalidPersonGender      = person_service.ErrInvalidGender
	ErrInvalidPersonNationality = person_service.ErrInvalidNationality
)

type (
	PersonServiceI person_service.ServiceI

	PersonCreatorI    person_service.CreatorI
	PersonDeleterI    person_service.DeleterI
	PersonUpdaterI    person_service.UpdaterI
	PersonGetterI     person_service.GetterI
	PersonListGetterI person_service.ListGetterI
)

func NewPersonService(
	log *slog.Logger,
	validator validation.ValidatorI,
	storage provider.PersonProviderI,
) PersonServiceI {
	return person_service.New(log, validator, storage)
}

type (
	PersonEnricherI person_enricher.EnricherI

	PersonByNameEnricherI person_enricher.ByNameEnricherI
)

func NewPersonEnricher(
	log *slog.Logger,
	personCreator person_service.CreatorI,
	ageProvider provider.AgeProviderI,
	genderProvider provider.GenderProviderI,
	nationalityProvider provider.NationalityProviderI,
) PersonEnricherI {
	return person_enricher.New(
		log, personCreator,
		ageProvider,
		genderProvider,
		nationalityProvider)
}

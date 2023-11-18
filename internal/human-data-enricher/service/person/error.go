package person_service

import (
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

var (
	ErrMissingID = ierror.NewInvalidArgument(
		"missing id", "MISSING_ID")
	ErrInvalidID = ierror.NewInvalidArgument(
		"invalid id", "INVALID_ID")

	ErrMissingName        = model.ErrMissingName
	ErrMissingSurname     = model.ErrMissingSurname
	ErrMissingAge         = model.ErrMissingAge
	ErrMissingGender      = model.ErrMissingGender
	ErrMissingNationality = model.ErrMissingNationality
	ErrInvalidName        = model.ErrInvalidName
	ErrInvalidSurname     = model.ErrInvalidSurname
	ErrInvalidMiddleName  = model.ErrInvalidMiddleName
	ErrInvalidAge         = model.ErrInvalidAge
	ErrInvalidGender      = model.ErrInvalidGender
	ErrInvalidNationality = model.ErrInvalidNationality
)

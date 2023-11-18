package ivalidator

import (
	"log/slog"

	pgvtor "github.com/go-playground/validator/v10"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person/validation"
)

type Validator struct {
	log      *slog.Logger
	validate *pgvtor.Validate
}

var _ validation.ValidatorI = (*Validator)(nil)

func NewValidator(
	log *slog.Logger,
	vtor *pgvtor.Validate,
) *Validator {
	return &Validator{
		log:      log,
		validate: vtor,
	}
}

//nolint:gochecknoglobals // TODO: use closures
package ivalidator

import (
	"context"
	"database/sql"
	"errors"

	pgvtor "github.com/go-playground/validator/v10"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	person_service "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/service/person"
)

type tag string

const (
	required tag = "required"
	alpha    tag = "alpha"
	uuid     tag = "uuid"
	gte      tag = "gte"
	lte      tag = "lte"
)

func matchErr(matcher map[tag]error, err error) error {
	if err == nil {
		return nil
	}
	var errs []error
	//nolint: errorlint // Safe to use with range statements
	for _, e := range err.(pgvtor.ValidationErrors) {
		t := tag(e.Tag())
		err, ok := matcher[t]
		if !ok {
			panic("violation with unknown tag")
		}

		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

var genderTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingGender,
	alpha:    person_service.ErrInvalidGender,
	uuid:     nil,
	gte:      nil,
	lte:      nil,
}

var nationalityTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingNationality,
	alpha:    person_service.ErrInvalidNationality,
	uuid:     nil,
	gte:      nil,
	lte:      nil,
}

var ageTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingAge,
	alpha:    person_service.ErrInvalidAge,
	uuid:     nil,
	gte:      person_service.ErrInvalidAge,
	lte:      person_service.ErrInvalidAge,
}

var nameTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingName,
	alpha:    person_service.ErrInvalidName,
	uuid:     nil,
	gte:      nil,
	lte:      nil,
}

var surnameTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingSurname,
	alpha:    person_service.ErrInvalidSurname,
	uuid:     nil,
	gte:      nil,
	lte:      nil,
}

var middleNameTagToErrMatcher = map[tag]error{
	required: nil,
	alpha:    person_service.ErrInvalidMiddleName,
	uuid:     nil,
	gte:      nil,
	lte:      nil,
}

var idTagToErrMatcher = map[tag]error{
	required: person_service.ErrMissingID,
	alpha:    nil,
	uuid:     person_service.ErrInvalidID,
	gte:      nil,
	lte:      nil,
}

func (v *Validator) Age(ctx context.Context, a int) error {
	return matchErr(
		ageTagToErrMatcher,
		v.validate.VarCtx(ctx, a, "required,gte=0,lte=130"),
	)
}

func (v *Validator) Gender(ctx context.Context, g string) error {
	return matchErr(
		genderTagToErrMatcher,
		v.validate.VarCtx(ctx, g, "required,alpha"),
	)
}

func (v *Validator) Nationality(ctx context.Context, n string) error {
	return matchErr(
		nationalityTagToErrMatcher,
		v.validate.VarCtx(ctx, n, "required,alpha"),
	)
}

func (v *Validator) Name(ctx context.Context, n string) error {
	return matchErr(
		nameTagToErrMatcher,
		v.validate.VarCtx(ctx, n, "required,alpha"),
	)
}

func (v *Validator) Surname(ctx context.Context, s string) error {
	return matchErr(
		surnameTagToErrMatcher,
		v.validate.VarCtx(ctx, s, "required,alpha"),
	)
}

func (v *Validator) MiddleName(ctx context.Context, m sql.NullString) error {
	return matchErr(
		middleNameTagToErrMatcher,
		v.validate.VarCtx(ctx, m.String, "omitempty,alpha"),
	)
}

func (v *Validator) ID(ctx context.Context, id model.PersonID) error {
	return matchErr(
		idTagToErrMatcher,
		v.validate.VarCtx(ctx, id, "required,uuid"),
	)
}

func (v *Validator) Person(ctx context.Context, p *model.Person) error {
	return errors.Join(
		v.Name(ctx, p.Name),
		v.Surname(ctx, p.Surname),
		v.MiddleName(ctx, p.MiddleName),
		v.Gender(ctx, p.Gender),
		v.Nationality(ctx, p.Nationality),
		v.Age(ctx, p.Age),
	)
}

package person

import (
	"errors"
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type (
	ValidatorFunc   func(*Person) error
	tagToErrMatcher map[tag]error
)

type field string

const (
	name  field = "Name"
	sname field = "Surname"
	mname field = "MiddleName"
)

type tag string

const (
	required tag = "required"
	alpha    tag = "alpha"
)

func NewValidator(
	log *slog.Logger,
	validate *validator.Validate,
) ValidatorFunc {
	nameMatcher := nameErrMatcher()
	snameMatcher := surnameErrMatcher()
	mnameMatcher := middleNameMatcher()

	matchers := map[field]tagToErrMatcher{
		name:  nameMatcher,
		sname: snameMatcher,
		mname: mnameMatcher,
	}

	return func(p *Person) error {
		var errs []error
		if err := validate.Struct(p); err != nil {
			//nolint: errorlint // safe to use with range statements
			for _, e := range err.(validator.ValidationErrors) {
				f := field(e.Field())
				matcher, ok := matchers[f]
				if !ok {
					log.Error("violation in unknown field", "field", f)
					continue
				}

				t := tag(e.Tag())
				err, ok := matcher[t]
				if !ok {
					log.Error("violation with unknown tag", "field", f, "tag", t)
					continue
				}

				errs = append(errs, err)
			}
		}
		return errors.Join(errs...)
	}
}

func nameErrMatcher() tagToErrMatcher {
	return map[tag]error{
		required: ErrMissingName,
		alpha:    ErrInvalidName,
	}
}

func surnameErrMatcher() tagToErrMatcher {
	return map[tag]error{
		required: ErrMissingSurname,
		alpha:    ErrInvalidSurname,
	}
}

func middleNameMatcher() tagToErrMatcher {
	return map[tag]error{
		required: nil,
		alpha:    ErrInvalidMiddleName,
	}
}

package ient

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
)

func (s *PersonStorage) Update(
	ctx context.Context,
	person *model.Person,
) (model.Person, error) {
	p := *person

	res, err := s.db.Person.
		UpdateOneID(person.ID()).
		SetName(person.Name).
		SetSurname(person.Surname).
		SetMiddleName(person.MiddleName).
		SetGender(person.Gender).
		SetNationality(person.Nationality).
		SetAge(person.Age).
		Save(ctx)
	switch {
	case err == nil:
		return toModel(res), nil

	case gen.IsNotFound(err):
		return p, provider.ErrPersonNotFound

	default:
		s.log.Error("bad attempt to update a person",
			"err", err,
			"res", res)
	}

	return p, err
}

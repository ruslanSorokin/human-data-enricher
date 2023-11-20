package ient

import (
	"context"

	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

func (s *PersonStorage) Create(
	ctx context.Context,
	person *model.Person,
) (model.Person, error) {
	var p model.Person
	res, err := s.db.Person.Create().
		SetID(person.ID()).
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

	case sqlgraph.IsUniqueConstraintError(err):
		return p, provider.ErrPersonAlreadyExists.
			Instantiate().
			WithDuplicateID(person.ID().String()).
			WithUniqueProperty("id")

	default:
		s.log.Error("bad attempt to create a person",
			"err", err,
			"person", person,
			"res", res)
	}

	return p, err
}

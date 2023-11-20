package ient

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/person"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/predicate"
)

func (s *PersonStorage) Get(
	ctx context.Context,
	id model.PersonID,
) (model.Person, error) {
	var p model.Person
	preds := []predicate.Person{
		person.DeletedAtIsNil(),
		person.ID(uuid.UUID(id)),
	}

	res, err := s.db.Person.
		Query().
		Where(preds...).
		Only(ctx)
	switch {
	case err == nil:
		return toModel(res), nil

	case gen.IsNotFound(err):
		return p, provider.ErrPersonNotFound

	default:
		s.log.Error("bad attempt to retrieve a person",
			"err", err,
			"id", id,
			"res", res)
	}

	return p, err
}

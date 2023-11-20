package ient

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
)

func (s *PersonStorage) Delete(ctx context.Context, id model.PersonID) error {
	err := s.db.Person.
		DeleteOneID(uuid.UUID(id)).
		Exec(ctx)

	switch {
	case err == nil:

	case gen.IsNotFound(err):
		return provider.ErrPersonNotFound

	default:
		s.log.Error("bad attempt to create a person",
			"err", err,
			"id", id)
	}

	return err
}

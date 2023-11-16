package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
)

type DeleterI interface {
	// Delete removes a Person object by "id".
	//
	// Known errors that may be returned:
	//
	// 	- ErrPersonNotFound
	//
	// 	- ErrMissingID
	//
	// 	- ErrInvalidID
	Delete(
		ctx context.Context,
		id model.PersonID,
	) error
}

func (s *PersonService) Delete(
	ctx context.Context,
	id model.PersonID,
) error {
	if err := s.vtor.ID(ctx, id); err != nil {
		return err
	}
	return s.storage.Delete(ctx, id)
}

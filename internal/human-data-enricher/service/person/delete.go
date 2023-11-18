package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
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
		if !ierror.As(err) {
			s.log.Error("bad attempt to create a person",
				"id", id)
		}
		return err
	}
	err := s.storage.Delete(ctx, id)
	switch {
	case err == nil || ierror.As(err):
	default:
		s.log.Error("bad attempt to delete a person",
			"err", err,
			"id", id)
	}
	return err
}

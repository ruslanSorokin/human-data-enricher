package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

type GetterI interface {
	// Get returns a Person object by "id".
	//
	// Known errors that may be returned:
	//
	// 	- ErrPersonNotFound
	//
	// 	- ErrMissingID
	//
	// 	- ErrInvalidID
	Get(
		ctx context.Context,
		id model.PersonID,
	) (model.Person, error)
}

func (s PersonService) Get(
	ctx context.Context,
	id model.PersonID,
) (model.Person, error) {
	var res model.Person

	if err := s.vtor.ID(ctx, id); err != nil {
		if !ierror.As(err) {
			s.log.Error("bad attempt to create a person",
				"id", id)
		}
		return res, err
	}
	res, err := s.storage.Get(ctx, id)
	switch {
	case err == nil || ierror.As(err):
	default:
		s.log.Error("bad attempt to retrieve a person",
			"err", err,
			"id", id)
	}
	return res, err
}

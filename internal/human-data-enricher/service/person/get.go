package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
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
		return res, err
	}
	res, err := s.storage.Get(ctx, id)
	if err != nil {
		return res, err
	}
	return res, nil
}

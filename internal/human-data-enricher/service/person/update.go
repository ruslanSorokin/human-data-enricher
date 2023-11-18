package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

type UpdaterI interface {
	// Update returns an updated Person object by "p.ID()".
	//
	// Known errors that may be returned:
	//
	// 	- ErrPersonNotFound
	//
	// 	- ErrMissingID
	//
	// 	- ErrInvalidID
	//
	// 	- ErrMissingName
	//
	// 	- ErrMissingSurname
	//
	// 	- ErrMissingAge
	//
	// 	- ErrMissingGender
	//
	// 	- ErrMissingNationality
	//
	// 	- ErrInvalidName
	//
	// 	- ErrInvalidSurname
	//
	// 	- ErrInvalidMiddleName
	//
	// 	- ErrInvalidAge
	//
	// 	- ErrInvalidGender
	//
	// 	- ErrInvalidNationality
	Update(
		ctx context.Context,
		p *model.Person,
	) (model.Person, error)
}

func (s *PersonService) Update(
	ctx context.Context,
	person *model.Person,
) (model.Person, error) {
	var res model.Person

	if err := s.vtor.Person(ctx, person); err != nil {
		if !ierror.As(err) {
			s.log.Error("bad attempt to create a person",
				"person", person)
		}
		return res, err
	}

	p, err := s.storage.Update(ctx, person)
	switch {
	case err == nil || ierror.As(err):
	default:
		s.log.Error("bad attempt to update a person",
			"err", err,
			"person", person)
	}
	return p, err
}

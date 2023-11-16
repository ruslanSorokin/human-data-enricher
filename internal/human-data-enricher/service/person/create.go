package person_service

import (
	"context"
	"errors"
	"strings"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

type CreatorI interface {
	// Create creates a new Person object. It validates opts, sanitizes all fields
	// and creates a Person object.
	//
	// Known errors that may be returned:
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
	Create(
		ctx context.Context,
		opts *model.PersonOptions,
	) (model.Person, error)
}

func (s *PersonService) Create(
	ctx context.Context,
	opts *model.PersonOptions,
) (model.Person, error) {
	opts = sanitazePersonOpts(opts)

	p := model.NewPerson(opts)

	if err := s.vtor.Person(ctx, &p); err != nil {
		s.log.Debug("bad attempt to create a person",
			"person", p,
			"opts", opts)
		return p, err
	}

	p, err := s.storage.Create(ctx, &p)
	switch {
	case err == nil:
	case errors.Is(err, provider.ErrPersonAlreadyExists):
	default:
		s.log.Error("bad attempt to insert a new person into the storage",
			"error", err,
			"person", p)
		return p, err
	}
	return p, nil
}

func sanitazePersonOpts(o *model.PersonOptions) *model.PersonOptions {
	o.Name = strings.TrimSpace(o.Name)
	o.Surname = strings.TrimSpace(o.Surname)
	o.MiddleName.String = strings.TrimSpace(o.MiddleName.String)
	return o
}

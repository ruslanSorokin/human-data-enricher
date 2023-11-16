package provider

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin"
)

var (
	// ErrInvalidPaginationRequest occurs if pagination request on person was.
	ErrInvalidPaginationRequest = ierror.NewInvalidArgument(
		"invalid pagination request on person", "INVALID_PAGINATION_REQUEST_ON_PERSON")

	// ErrPersonAlreadyExists occurs if person with such "ID" already exists.
	ErrPersonAlreadyExists = ierror.NewAlreadyExists(
		"person already exists", "PERSON_ALREADY_EXISTS")

	// ErrPersonNotFound occurs if person with such "ID" not found.
	ErrPersonNotFound = ierror.NewNotFound(
		"person not found", "PERSON_NOT_FOUND")
)

type PersonProviderI interface {
	PersonCreatorI
	PersonUpdaterI
	PersonDeleterI
	PersonGetterI

	PersonListGetterI
}

type (
	PersonCreatorI interface {
		Create(
			ctx context.Context,
			p *model.Person,
		) (model.Person, error)
	}

	PersonGetterI interface {
		Get(
			ctx context.Context,
			id model.PersonID,
		) (model.Person, error)
	}

	PersonUpdaterI interface {
		Update(
			ctx context.Context,
			p *model.Person,
		) (model.Person, error)
	}

	PersonDeleterI interface {
		Delete(
			ctx context.Context,
			id model.PersonID,
		) error
	}

	PersonListGetterI interface {
		GetList(
			ctx context.Context,
			pagTkn pagin.ReqToken,
		) ([]*model.Person, pagin.RespToken, error)
	}
)

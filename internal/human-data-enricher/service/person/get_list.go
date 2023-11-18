package person_service

import (
	"context"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin"
)

type ListGetterI interface {
	// GetList returns a list of Persons object by "pgTkn".
	//
	// Known errors that may be returned:
	//
	// 	- ErrInvalidPaginationRequest
	GetList(
		ctx context.Context,
		pgTkn pagin.ReqToken,
	) ([]*model.Person, pagin.RespToken, error)
}

func (s *PersonService) GetList(
	ctx context.Context,
	pagTkn pagin.ReqToken,
) ([]*model.Person, pagin.RespToken, error) {
	return s.storage.GetList(ctx, pagTkn)
}

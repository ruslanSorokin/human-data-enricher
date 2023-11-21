package ient

import (
	"context"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/person"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen/predicate"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin"
)

const MinLimit = 0

//nolint:gochecknoglobals // TODO: use closures
var fieldMatcher = map[string]string{
	"created_at":  person.FieldCreatedAt,
	"updated_at":  person.FieldUpdatedAt,
	"name":        person.FieldName,
	"surname":     person.FieldSurname,
	"middle_name": person.FieldMiddleName,
	"gender":      person.FieldGender,
	"nationality": person.FieldNationality,
	"age":         person.FieldAge,
}

func (s *PersonStorage) GetList(
	ctx context.Context,
	reqTkn pagin.ReqToken,
) ([]*model.Person, pagin.RespToken, error) {
	respTkn := pagin.NewRespToken(reqTkn)

	if reqTkn.Limit <= MinLimit {
		return nil, respTkn, provider.ErrInvalidPaginationRequest
	}

	isFirstPage := reqTkn.Cursor.IsEmpty()
	preds := []predicate.Person{person.DeletedAtIsNil()}
	if !isFirstPage {
		p, err := cursorPredicate(reqTkn)
		if err != nil {
			return nil, respTkn, err
		}
		preds = append(preds, p)
	}

	if !isFirstPage && reqTkn.HasFieldOrdering() {
		p, err := fieldPredicates(reqTkn)
		if err != nil {
			return nil, respTkn, err
		}
		preds = append(preds, p)
	}

	isDesc := reqTkn.Order.IsDesc()
	ordering := make([]person.OrderOption, 0, len(reqTkn.OrderBy))
	if reqTkn.HasFieldOrdering() {
		for _, ord := range reqTkn.OrderBy {
			isFieldDesc := ord.Order == pagin.Desc
			o, err := fieldOrdering(ord.Field, isFieldDesc)
			if err != nil {
				return nil, respTkn, err
			}
			ordering = append(ordering, o)
		}
	}
	// WARN: Must be the last one in the ordering slice.
	ordering = append(ordering, idOrdering(isDesc))

	dbLimit := reqTkn.Limit + 1 // limit is 2 or more
	pp, err := s.db.Person.
		Query().
		Where(preds...).
		Order(ordering...).
		Limit(dbLimit).
		All(ctx)
	if err != nil {
		return nil, respTkn, err
	}

	if len(pp) == dbLimit {
		// We have the next page(at least 1 element, because we've fetched (n + 1)).
		// We should return "pp" without the last element.
		respTkn.HasNextPage = true

		if !isFirstPage {
			lhs := pp[0]
			respTkn.NextPageCursor.Right = lhs.ID.String()
		}

		rhs := pp[len(pp)-2]
		respTkn.NextPageCursor.Left = rhs.ID.String()
	}

	if len(pp) < dbLimit {
		// We don't have the next page, because we've fetched less than (n + 1)
		// We should return the whole "pp"
		respTkn.HasNextPage = false
	}

	results := make([]*model.Person, 0, len(pp))
	for i := 0; i < len(pp) && i < dbLimit-1; i++ {
		results = append(results, toModelP(pp[i]))
	}

	return results, respTkn, nil
}

func cursorPredicate(tkn pagin.ReqToken) (predicate.Person, error) {
	switch {
	case tkn.Cursor.IsLeft(), tkn.Cursor.IsRight():
		id, err := uuid.FromString(tkn.Cursor.Cursor())
		if err != nil {
			return nil, err
		}
		return idPredicate(id, tkn.Cursor.IsRight()), nil

	case tkn.Cursor.IsBounded():
		lID, err := uuid.FromString(tkn.Cursor.Left)
		if err != nil {
			return nil, err
		}

		rID, err := uuid.FromString(tkn.Cursor.Right)
		if err != nil {
			return nil, err
		}

		return person.And(
			idPredicate(lID, false),
			idPredicate(rID, true),
		), nil

	default:
		return nil, provider.ErrInvalidPaginationRequest
	}
}

func fieldOrdering(field string, isDesc bool) (person.OrderOption, error) {
	f, ok := fieldMatcher[field]
	if !ok {
		return nil, provider.ErrInvalidPaginationRequest
	}

	if isDesc {
		return gen.Desc(f), nil
	}
	return gen.Asc(f), nil
}

func fieldPredicates(tkn pagin.ReqToken) (predicate.Person, error) {
	var ret []predicate.Person
	for _, ord := range tkn.OrderBy {
		f, ok := fieldMatcher[ord.Field]
		if !ok {
			return nil, provider.ErrInvalidPaginationRequest
		}

		switch {
		case tkn.Cursor.IsLeft(), tkn.Cursor.IsRight():
			id, err := uuid.FromString(tkn.Cursor.Cursor())
			if err != nil {
				return nil, err
			}
			return idPredicate(id, tkn.Cursor.IsRight()), nil

		case tkn.Cursor.IsBounded():
			l, err := fieldPredicate(f, ord.Cursor.Left, false)
			if err != nil {
				return nil, err
			}
			r, err := fieldPredicate(f, ord.Cursor.Right, true)
			if err != nil {
				return nil, err
			}

			ret = append(ret, person.And(l, r))

		default:
			return nil, provider.ErrInvalidPaginationRequest
		}
	}

	return person.And(ret...), nil
}

func idOrdering(isDesc bool) person.OrderOption {
	if isDesc {
		return gen.Desc(person.FieldID)
	}
	return gen.Asc(person.FieldID)
}

func idPredicate(id uuid.UUID, isRight bool) predicate.Person {
	if isRight {
		return person.IDLT(id)
	}
	return person.IDGT(id)
}

func fieldPredicate(f string, v any, isRight bool) (predicate.Person, error) {
	f, ok := fieldMatcher[f]
	if !ok {
		return nil, provider.ErrInvalidPaginationRequest
	}

	if isRight {
		return predicate.Person(sql.FieldLT(f, v)), nil
	}
	return predicate.Person(sql.FieldGT(f, v)), nil
}

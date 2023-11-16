package pagin

import "github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin/cursor"

type order string

const (
	Asc  order = "ASC"
	Desc order = "DESC"
)

func (o *order) IsDesc() bool {
	return *o == Desc
}

func (o *order) IsAsc() bool {
	return *o == Desc
}

type OrderByCell struct {
	Field  string        `json:"field,omitempty"`
	Order  order         `json:"order,omitempty"`
	Cursor cursor.Cursor `json:"cursor"`
}

type ReqToken struct {
	Cursor cursor.Cursor `json:"cursor"`
	Order  order         `json:"order"`

	OrderBy []OrderByCell `json:"order_by,omitempty"`
	Limit   int           `json:"limit"`
}

func NewReqToken(opts ...ReqTokenOption) ReqToken {
	var t ReqToken

	for _, o := range opts {
		o(&t)
	}

	return t
}

func (t *ReqToken) HasCursor() bool {
	// No need to check cursors of field ordering, since we won't be able to
	// paginate without cursor on ID.
	return !t.Cursor.IsEmpty()
}

func (t *ReqToken) HasFieldOrdering() bool {
	return len(t.OrderBy) != 0
}

type ReqTokenOption = func(*ReqToken)

func WithCursor(c cursor.Cursor) ReqTokenOption {
	return func(t *ReqToken) {
		t.Cursor = c
	}
}

func WithLimit(l int) ReqTokenOption {
	return func(t *ReqToken) {
		t.Limit = l
	}
}

type OrderByOptions struct {
	Field  string
	Cursor cursor.Cursor
	IsDesc bool
}

func (opt *OrderByOptions) toCell() OrderByCell {
	return OrderByCell{
		Field: opt.Field,
		Order: func() order {
			if opt.IsDesc {
				return Desc
			}
			return Asc
		}(),
		Cursor: opt.Cursor,
	}
}

func WithOrderBy(opts ...OrderByOptions) ReqTokenOption {
	return func(t *ReqToken) {
		for _, opt := range opts {
			t.OrderBy = append(t.OrderBy, opt.toCell())
		}
	}
}

func WithDescOrder() ReqTokenOption {
	return func(t *ReqToken) {
		t.Order = Desc
	}
}

func WithAscOrder() ReqTokenOption {
	return func(t *ReqToken) {
		t.Order = Asc
	}
}

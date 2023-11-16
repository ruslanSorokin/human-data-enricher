package pagin

import "github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin/cursor"

type RespToken struct {
	OrderBy        []OrderByCell `json:"order_by,omitempty"`
	PrevPageCursor cursor.Cursor `json:"prev_page_cursor"`
	NextPageCursor cursor.Cursor `json:"next_page_cursor"`

	ReqToken ReqToken `json:"request_token"`

	HasNextPage bool `json:"has_next_page"`
}

func (t *RespToken) NextPageTkn() ReqToken {
	return ReqToken{
		Cursor:  t.NextPageCursor,
		OrderBy: t.OrderBy,
		Order:   t.ReqToken.Order,
		Limit:   t.ReqToken.Limit,
	}
}

func (t *RespToken) PrevPageTkn() ReqToken {
	return ReqToken{
		Cursor:  t.PrevPageCursor,
		OrderBy: t.OrderBy,
		Order:   t.ReqToken.Order,
		Limit:   t.ReqToken.Limit,
	}
}

func NewRespToken(t ReqToken) RespToken {
	return RespToken{
		ReqToken:       t,
		PrevPageCursor: cursor.Nil,
		NextPageCursor: cursor.Nil,
		OrderBy:        []OrderByCell{},
		HasNextPage:    false,
	}
}

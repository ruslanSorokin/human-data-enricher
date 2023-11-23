package test_person

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util/pagin/cursor"
)

func (s PSuite) TestGetList_Playground() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	ppInitial := aListOfPersons()

	for _, p := range ppInitial {
		_, err := repo.Create(ctx, p)
		assert.NoError(err)
	}

	reqTkn := pagin.NewReqToken(
		pagin.WithLimit(50),
		pagin.WithAscOrder(),
		pagin.WithOrderBy(
			pagin.OrderByOptions{Field: "age", Cursor: cursor.Nil, IsDesc: false},
		),
	)
	pp, _, err := repo.GetList(ctx, reqTkn)

	for _, p := range pp {
		t.Log(p.Age)
	}

	require.NoError(err)
}

func (s PSuite) TestGetList_Asc_Successful() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	testCases := []struct {
		lim int
	}{
		{lim: 1},
		{lim: 2},
		{lim: 3},
		{lim: 4},
		{lim: 5},
		{lim: 6},
		{lim: 7},
		{lim: 8},
	}

	ppInitial := aListOfPersons()

	for _, p := range ppInitial {
		_, err := repo.Create(ctx, p)
		assert.NoError(err)
	}

	for _, tc := range testCases {
		t.Log(tc.lim)
		reqTkn := pagin.NewReqToken(
			pagin.WithLimit(tc.lim),
		)
		ppGot, respTkn, err := repo.GetList(ctx, reqTkn)
		require.NoError(err)
		offset := len(ppGot)
		for i := 0; i < len(ppGot); i++ {
			require.EqualExportedValues(*ppInitial[i], *ppGot[i]) // from 0 to lim
		}
		ppGot, respTkn, err = repo.GetList(ctx, respTkn.NextPageTkn())
		require.NoError(err)
		if respTkn.HasNextPage {
			for i := 0; i < len(ppGot); i++ {
				require.EqualExportedValues(
					*ppInitial[i+offset],
					*ppGot[i],
				) // len(firstBatch) to n
			}
		}
	}
}

func (s PSuite) TestGetList_Desc_Successful() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	testCases := []struct {
		lim int
	}{
		{lim: 1},
		{lim: 2},
		{lim: 3},
		{lim: 4},
		{lim: 5},
		{lim: 6},
		{lim: 7},
		{lim: 8},
	}

	ppInitial := aListOfPersons()

	for _, p := range ppInitial {
		_, err := repo.Create(ctx, p)
		assert.NoError(err)
	}

	for i, j := 0, len(ppInitial)-1; i < j; i, j = i+1, j-1 {
		ppInitial[i], ppInitial[j] = ppInitial[j], ppInitial[i] // reverse
	}

	for _, tc := range testCases {
		t.Log(tc.lim)
		reqTkn := pagin.NewReqToken(
			pagin.WithLimit(tc.lim),
			pagin.WithDescOrder(),
		)
		ppGot, respTkn, err := repo.GetList(ctx, reqTkn)
		require.NoError(err)
		offset := len(ppGot)
		for i := 0; i < len(ppGot); i++ {
			require.EqualExportedValues(*ppInitial[i], *ppGot[i]) // from 0 to lim
		}
		ppGot, respTkn, err = repo.GetList(ctx, respTkn.NextPageTkn())
		require.NoError(err)
		if respTkn.HasNextPage {
			for i := 0; i < len(ppGot); i++ {
				require.EqualExportedValues(
					*ppInitial[i+offset],
					*ppGot[i],
				) // len(firstBatch) to n
			}
		}
	}
}

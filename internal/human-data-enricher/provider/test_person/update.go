package test_person

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

func (s PSuite) TestUpdate_Successful() {
	t := s.T()
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()
	repo := s.Provider

	p := aPerson()

	_, err := repo.Create(ctx, &p)
	assert.NoError(err,
		"should insert without error")

	_, err = repo.Update(ctx, &p)
	require.NoError(err,
		"must update without error")
}

func (s PSuite) TestUpdate_NotFound() {
	t := s.T()
	require := require.New(t)
	ctx := context.Background()
	repo := s.Provider

	p := aPerson()
	_, err := repo.Update(ctx, &p)

	wantErr := provider.ErrPersonNotFound
	require.ErrorIs(err, wantErr,
		"must return %w", wantErr)
}

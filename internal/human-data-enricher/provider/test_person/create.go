package test_person

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

func (s PSuite) TestCreate_Successful() {
	t := s.T()
	repo := s.Provider
	require := require.New(t)
	ctx := context.Background()

	p := aPerson()

	_, err := repo.Create(ctx, &p)
	require.NoError(err,
		"must insert without error")
}

func (s PSuite) TestCreate_AlreadyExists() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	p := aPerson()

	_, err := repo.Create(ctx, &p)
	assert.NoError(err,
		"should insert without error")

	_, err = repo.Create(ctx, &p)
	wantErr := provider.ErrPersonAlreadyExists
	require.ErrorIs(err, wantErr,
		"must return %w", wantErr)
}

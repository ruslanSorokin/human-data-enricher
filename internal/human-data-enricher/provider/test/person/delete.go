package test_person

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

func (s PSuite) TestDelete_Successful() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	p := aPerson()

	_, err := repo.Create(ctx, &p)
	assert.NoError(err,
		"should insert without error")

	err = repo.Delete(ctx, model.PersonID(p.ID()))
	require.NoError(err,
		"must return error not found")
}

func (s PSuite) TestDelete_NotFound() {
	t := s.T()
	repo := s.Provider
	require := require.New(t)
	ctx := context.Background()

	err := repo.Delete(ctx, model.PersonID(uuid.Must(uuid.NewV7())))
	wantErr := provider.ErrPersonNotFound
	require.ErrorIs(err, wantErr,
		"must return %w", wantErr)
}

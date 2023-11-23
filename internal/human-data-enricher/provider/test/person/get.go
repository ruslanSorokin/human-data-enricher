package test_person

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/model"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
)

func (s PSuite) TestGet_Successful() {
	t := s.T()
	repo := s.Provider
	assert := assert.New(t)
	require := require.New(t)
	ctx := context.Background()

	p := aPerson()

	want, err := repo.Create(ctx, &p)
	assert.NoError(err,
		"should insert without error")

	got, err := repo.Get(ctx, model.PersonID(p.ID()))
	require.NoError(err,
		"must return person without error")

	require.EqualExportedValues(want, got,
		"must be equals")
}

func (s PSuite) TestGet_NotFound() {
	t := s.T()
	repo := s.Provider
	require := require.New(t)
	ctx := context.Background()

	id := uuid.Must(uuid.NewV7())

	_, err := repo.Get(ctx, model.PersonID(id))
	wantErr := provider.ErrPersonNotFound
	require.ErrorIs(err, wantErr,
		"must return %w", wantErr)
}

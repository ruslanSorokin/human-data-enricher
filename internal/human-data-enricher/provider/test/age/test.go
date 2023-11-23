package test_age

import (
	"context"

	"github.com/stretchr/testify/require"
)

const (
	minAge = 0
	maxAge = 130
)

func (s ProviderSuite) TestFetch_Successful() {
	t := s.T()
	gw := s.Provider
	require := require.New(t)
	ctx := context.Background()

	a, err := gw.AgeByName(ctx, "Dmitry")
	require.NoError(err)
	require.True(minAge <= a && a <= maxAge)
}

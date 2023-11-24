package test_gender

import (
	"context"

	"github.com/stretchr/testify/require"
)

func (s ProviderSuite) TestFetch_Successful() {
	t := s.T()
	gw := s.Provider
	require := require.New(t)
	ctx := context.Background()

	g, err := gw.GenderByName(ctx, "Dmitry")
	require.NoError(err)
	require.NotEmpty(g)
}

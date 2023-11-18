package ierror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

func TestInstantiateAlreadyExists(t *testing.T) {
	require := require.New(t)

	userStatic := ierror.NewAlreadyExists(
		"user with given login already exists",
		"USER_ALREADY_EXISTS",
	)
	dupIDUser := "73546234"

	userDynamic := userStatic.Instantiate().WithDuplicateID(dupIDUser)

	dup, ok := userDynamic.DuplicateID()

	require.True(ok)
	require.Equal(dup, dupIDUser)

	require.False(errors.Is(userStatic, userDynamic))

	productStatic := ierror.NewAlreadyExists(
		"product with given name already exists",
		"PRODUCT_ALREADY_EXISTS",
	)

	dupIDProduct := "6456235"
	productDynamic := productStatic.Instantiate().WithDuplicateID(dupIDProduct)

	dup, ok = productDynamic.DuplicateID()

	require.True(ok)
	require.Equal(dup, dupIDProduct)

	require.True(errors.Is(userDynamic, userStatic))
	require.False(errors.Is(userStatic, userDynamic))

	require.True(errors.Is(productDynamic, productStatic))
	require.False(errors.Is(productStatic, productDynamic))

	require.False(errors.Is(productStatic, productDynamic))

	require.False(errors.Is(productStatic, userDynamic))
	require.False(errors.Is(userStatic, productDynamic))

	require.False(errors.Is(userStatic, errors.New("random error here")))
}

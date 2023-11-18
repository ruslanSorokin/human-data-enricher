package ierror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/ierror"
)

func TestAPIError_WithEnum(t *testing.T) {
	require := require.New(t)

	userStatic := ierror.New(
		"user with given login already exists",
		codes.AlreadyExists, http.StatusConflict,
		"USER_ALREADY_EXISTS")

	userDynamic := userStatic.Instantiate().WithEnum("73546234")

	productStatic := ierror.New(
		"product with given name already exists",
		codes.AlreadyExists, http.StatusConflict,
		"PRODUCT_ALREADY_EXISTS")

	productDynamic := productStatic.Instantiate().WithEnum("6456235")

	require.True(errors.Is(userDynamic, userStatic))
	require.False(errors.Is(userStatic, userDynamic))

	require.True(errors.Is(productDynamic, productStatic))
	require.False(errors.Is(productStatic, productDynamic))

	require.False(errors.Is(productStatic, userDynamic))
	require.False(errors.Is(userStatic, productDynamic))

	require.False(errors.Is(userStatic, errors.New("random error here")))
}

func TestAPIError_WithGRPCStCode(t *testing.T) {
	require := require.New(t)

	userStatic := ierror.New(
		"user with given login already exists",
		codes.AlreadyExists, http.StatusConflict,
		"USER_ALREADY_EXISTS")

	userDynamic := userStatic.Instantiate().WithGRPCStCode(codes.DataLoss)

	productStatic := ierror.New(
		"product with given name already exists",
		codes.AlreadyExists, http.StatusConflict,
		"PRODUCT_ALREADY_EXISTS")

	productDynamic := productStatic.Instantiate().WithGRPCStCode(codes.DataLoss)

	require.True(errors.Is(userDynamic, userStatic))
	require.False(errors.Is(userStatic, userDynamic))

	require.True(errors.Is(productDynamic, productStatic))
	require.False(errors.Is(productStatic, productDynamic))

	require.False(errors.Is(productStatic, userDynamic))
	require.False(errors.Is(userStatic, productDynamic))

	require.False(errors.Is(userStatic, errors.New("random error here")))
}

func TestAPIError_WithHTTPStCode(t *testing.T) {
	require := require.New(t)

	userStatic := ierror.New(
		"user with given login already exists",
		codes.AlreadyExists, http.StatusConflict,
		"USER_ALREADY_EXISTS")

	userDynamic := userStatic.Instantiate().
		WithHTTPStCode(http.StatusExpectationFailed)

	productStatic := ierror.New(
		"product with given name already exists",
		codes.AlreadyExists, http.StatusConflict,
		"PRODUCT_ALREADY_EXISTS")

	productDynamic := productStatic.Instantiate().
		WithHTTPStCode(http.StatusExpectationFailed)

	require.True(errors.Is(userDynamic, userStatic))
	require.False(errors.Is(userStatic, userDynamic))

	require.True(errors.Is(productDynamic, productStatic))
	require.False(errors.Is(productStatic, productDynamic))

	require.False(errors.Is(productStatic, userDynamic))
	require.False(errors.Is(userStatic, productDynamic))

	require.False(errors.Is(userStatic, errors.New("random error here")))
}

func TestAPIError_Getters(t *testing.T) {
	require := require.New(t)

	message := "message"
	grpc := codes.OK
	http := http.StatusOK
	enum := "enum"

	err := ierror.New(message, grpc, http, enum)

	require.Equal(err.ToEnum(), enum)
	require.Equal(err.ToGRPC(), grpc)
	require.Equal(err.ToHTTP(), http)
	require.Equal(err.Error(), message)
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"openapi"
	"openapi/internal/utils"
	"openapi/models/components"
	"testing"
)

func TestStore_GetInventory(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getInventory")

	s := openapi.New(
		openapi.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		openapi.WithClient(testHTTPClient),
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Store.GetInventory(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, map[string]int{
		"key":  373538,
		"key1": 961069,
	}, res.Object)

}

func TestStore_PlaceOrder(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("placeOrder")

	s := openapi.New(
		openapi.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		openapi.WithClient(testHTTPClient),
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Store.PlaceOrder(ctx, &components.Order{
		ID:       openapi.Int64(10),
		PetID:    openapi.Int64(198772),
		Quantity: openapi.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Order)
	assert.Equal(t, &components.Order{
		ID:       openapi.Int64(10),
		PetID:    openapi.Int64(198772),
		Quantity: openapi.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)

}

func TestStore_GetOrderByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getOrderById")

	s := openapi.New(
		openapi.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		openapi.WithClient(testHTTPClient),
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Store.GetOrderByID(ctx, 614993)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Order)
	assert.Equal(t, &components.Order{
		ID:       openapi.Int64(10),
		PetID:    openapi.Int64(198772),
		Quantity: openapi.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)

}

func TestStore_DeleteOrder(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("deleteOrder")

	s := openapi.New(
		openapi.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		openapi.WithClient(testHTTPClient),
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Store.DeleteOrder(ctx, 127902)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Order)
	assert.Equal(t, &components.Order{
		ID:       openapi.Int64(10),
		PetID:    openapi.Int64(198772),
		Quantity: openapi.Int(7),
		Status:   components.OrderStatusApproved.ToPointer(),
	}, res.Order)

}

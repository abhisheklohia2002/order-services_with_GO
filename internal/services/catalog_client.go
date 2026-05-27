package services

import (
	"context"

	ordersModel "example.com/m/v4/internal/models/orders"
)

type CatalogClient interface {
	GetProductByID(ctx context.Context, productID uint) (*ordersModel.ProductSnapshot, error)
}

package storage

import (
	"context"
)

type Storage interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductsByCategory(ctx context.Context, category string) ([]Product, error)
	GetProductsUnderPrice(ctx context.Context, priceUnder int) ([]Product, error)
	GetProductsByCategoryUnderPrice(ctx context.Context, category string, priceUnder int) ([]Product, error)
}

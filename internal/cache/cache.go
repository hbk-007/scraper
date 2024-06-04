package cache

import (
	"context"
	"scraper/internal/domain/entities"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	db *redis.Client
}

type ICache interface {
	GetProduct(ctx context.Context, productTitle string) *entities.Product
	UpdateProduct(ctx context.Context, product entities.Product) error
}

func NewCache(redisCli *redis.Client) ICache {
	return &Cache{
		db: redisCli,
	}
}

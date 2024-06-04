package cache

import (
	"context"
	"encoding/json"
	"log"
	"scraper/internal/domain/entities"
	"time"

	"github.com/go-redis/redis/v8"
)

func (c *Cache) GetProduct(ctx context.Context, productTitle string) *entities.Product {
	_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	data, err := c.db.Get(_ctx, productTitle).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Printf("Error getting product from cache: %v", err)
		return nil
	}

	var product entities.Product
	err = json.Unmarshal([]byte(data), &product)
	if err != nil {
		log.Printf("Error unmarshalling product data: %v", err)
		return nil
	}
	return &product
}

func (c *Cache) UpdateProduct(ctx context.Context, product entities.Product) error {
	_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return c.db.Set(_ctx, product.ProductTitle, data, 0).Err()
}

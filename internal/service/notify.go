package service

import (
	"context"
	"log"
)

func (s *Service) Notify(ctx context.Context, count int) error {
	log.Printf("Scraped %d products and updated the database.", count)
	return nil
}

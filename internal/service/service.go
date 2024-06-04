package service

import (
	"context"
	"scraper/internal/cache"
	_dentalstall "scraper/internal/clients/dentalstall"
	"scraper/internal/domain/entities"
	"scraper/internal/domain/request"
	"scraper/internal/repo"
)

type Service struct {
	DentalStallClient _dentalstall.IDentalStallClient
	Cache             cache.ICache
	Repo              repo.IScrapeRepo
}

type IClient interface {
	Scrape(ctx context.Context, request *request.ScrapeRequest) ([]entities.Product, error)
	Notify(ctx context.Context, count int) error
}

func NewService(dentalClient _dentalstall.IDentalStallClient, cache cache.ICache, repo repo.IScrapeRepo) IClient {
	return &Service{
		DentalStallClient: dentalClient,
		Cache:             cache,
		Repo:              repo,
	}
}

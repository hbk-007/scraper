package service

import (
	"context"
	"scraper/internal/domain/entities"
	"scraper/internal/domain/request"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (s *Service) Scrape(ctx context.Context, request *request.ScrapeRequest) ([]entities.Product, error) {
	allProducts := []entities.Product{}
	for page := 1; page <= request.PageLimit; page++ {
		products, err := s.scrapePage(page, request.Proxy)
		if err != nil {
			return nil, err
		}
		allProducts = append(allProducts, products...)
	}
	for _, product := range allProducts {
		cachedProduct := s.Cache.GetProduct(ctx, product.ProductTitle)
		if cachedProduct == nil || cachedProduct.ProductPrice != product.ProductPrice {
			s.Repo.SaveProduct(product)
			s.Cache.UpdateProduct(ctx, product)
		}
	}
	s.Notify(ctx, len(allProducts))
	return allProducts, nil
}

func (s *Service) scrapePage(pageNumber int, proxy string) ([]entities.Product, error) {
	resp, err := s.DentalStallClient.GetPageData(pageNumber, proxy)
	if err != nil {
		return nil, err
	}
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp)
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	doc.Find(".products li.product").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".woo-loop-product__title a").Text()
		priceStr := s.Find(".mf-product-price-box .price ins .woocommerce-Price-amount").Text()
		priceStr = strings.ReplaceAll(priceStr, "₹", "")
		priceStr = strings.ReplaceAll(priceStr, ",", "")
		price, _ := strconv.ParseFloat(priceStr, 64)
		imgURL, _ := s.Find(".mf-product-thumbnail img").Attr("src")

		product := entities.Product{ProductTitle: title, ProductPrice: price, PathToImage: imgURL}
		products = append(products, product)
	})
	return products, nil
}

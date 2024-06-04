package repo

import (
	"encoding/json"
	"os"
	"scraper/internal/domain/entities"
)

type Repository struct {
	filePath string
}

type IScrapeRepo interface {
	SaveProduct(product entities.Product) error
}

func NewRepository(filePath string) IScrapeRepo {
	return &Repository{filePath: filePath}
}

func (repo *Repository) loadData() ([]entities.Product, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []entities.Product{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var data []entities.Product
	err = json.NewDecoder(file).Decode(&data)
	return data, err
}

func (repo *Repository) saveData(data []entities.Product) error {
	file, err := os.Create(repo.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

package repo

import (
	"scraper/internal/domain/entities"
)

func (r *Repository) SaveProduct(product entities.Product) error {
	data, err := r.loadData()
	if err != nil {
		return err
	}

	data = append(data, product)
	return r.saveData(data)
}

package temp

import (
	"errors"

	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/port"
)

type tempProductRepository struct {
	localData map[string]model.Product
}

func NewTempProductRepository() port.ProductRepository {
	return &tempProductRepository{
		localData: make(map[string]model.Product),
	}
}

func (repo *tempProductRepository) Create(product model.Product) error {
	_, isExist := repo.localData[product.Id]
	if isExist {
		return errors.New("product already exist")
	}
	repo.localData[product.Id] = product
	return nil
}

func (repo *tempProductRepository) GetById(id string) (*model.Product, error) {
	product, isExist := repo.localData[id]
	if !isExist {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

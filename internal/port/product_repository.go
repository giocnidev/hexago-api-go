package port

import "github.com/giocnidev/api-hexagonal/internal/core/model"

type ProductRepository interface {
	Create(model.Product) error
	GetById(string) (*model.Product, error)
}

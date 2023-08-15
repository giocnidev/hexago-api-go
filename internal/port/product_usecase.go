package port

import "github.com/giocnidev/api-hexagonal/internal/core/model"

type CreateProductUseCase interface {
	Execute(model.Product) (*model.Product, *model.Error)
}

type GetProductById interface {
	Execute(string) (*model.Product, *model.Error)
}

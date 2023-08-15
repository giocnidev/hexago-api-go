package usecase

import (
	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/port"
)

type getProductByIdUseCase struct {
	productRepository port.ProductRepository
}

func NewGetProductByIdUseCase(productRepository port.ProductRepository) port.GetProductById {
	return &getProductByIdUseCase{
		productRepository: productRepository,
	}
}

func (useCase *getProductByIdUseCase) Execute(productId string) (*model.Product, *model.Error) {
	product, err := useCase.productRepository.GetById(productId)
	if err != nil {
		return nil, &model.Error{
			Code:    "",
			Message: err.Error(),
		}
	}
	return product, nil
}

package usecase

import (
	"errors"
	"strings"

	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/port"
)

type createProductUseCase struct {
	productRepository port.ProductRepository
}

func NewCreateProductUseCase(productRepository port.ProductRepository) port.CreateProductUseCase {
	return &createProductUseCase{
		productRepository: productRepository,
	}
}

func (useCase *createProductUseCase) Execute(product model.Product) (*model.Product, *model.Error) {
	if err := validateProduct(product); err != nil {
		return nil, &model.Error{
			Code:    "",
			Message: err.Error(),
		}
	}

	err := useCase.productRepository.Create(product)
	if err != nil {
		return nil, &model.Error{
			Code:    "",
			Message: err.Error(),
		}
	}

	return &product, nil
}

func validateProduct(product model.Product) error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("the product name is not correct.")
	}
	if product.Price < 0 {
		return errors.New("the product price is not correct.")
	}
	return nil
}

package core

import "github.com/giocnidev/api-hexagonal/internal/port"

type Application struct {
	CreateProductUseCase  port.CreateProductUseCase
	GetProductByIdUseCase port.GetProductById
}

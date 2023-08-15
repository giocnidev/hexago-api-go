package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giocnidev/api-hexagonal/internal/adapter/http"
	"github.com/giocnidev/api-hexagonal/internal/adapter/repository/temp"
	"github.com/giocnidev/api-hexagonal/internal/core"
	"github.com/giocnidev/api-hexagonal/internal/core/usecase"
)

func main() {
	app := gin.Default()

	productRepository := temp.NewTempProductRepository()
	application := core.Application{
		CreateProductUseCase:  usecase.NewCreateProductUseCase(productRepository),
		GetProductByIdUseCase: usecase.NewGetProductByIdUseCase(productRepository),
	}
	apiController := http.NewApiController(app, application)
	apiController.BindRoutes()
	app.Run()
}

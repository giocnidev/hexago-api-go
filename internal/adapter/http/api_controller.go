package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giocnidev/api-hexagonal/internal/core"
	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/port"
)

type apiController struct {
	ginEngine   *gin.Engine
	application core.Application
}

func NewApiController(ginEngine *gin.Engine, application core.Application) port.ApiController {
	return &apiController{
		ginEngine:   ginEngine,
		application: application,
	}
}

func (ctrl *apiController) BindRoutes() {
	api := ctrl.ginEngine.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/products", ctrl.createProductHandler)
		v1.GET("/products/:productId", ctrl.getProductByIdHandler)
	}
}

func (ctrl *apiController) createProductHandler(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
		})
		return
	}
	result, err := ctrl.application.CreateProductUseCase.Execute(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *apiController) getProductByIdHandler(ctx *gin.Context) {
	productId, ok := ctx.Params.Get("productId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Message: "product id is required",
		})
		return
	}
	product, err := ctrl.application.GetProductByIdUseCase.Execute(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

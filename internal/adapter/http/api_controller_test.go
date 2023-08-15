package http_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/giocnidev/api-hexagonal/internal/adapter/http"
	"github.com/giocnidev/api-hexagonal/internal/core"
	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/mock"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestBindRoutes(t *testing.T) {
	app := gin.Default()
	apiController := http.NewApiController(app, core.Application{})
	apiController.BindRoutes()
	routes := app.Routes()
	if app.Handlers == nil {
		t.Fatal("errors handlers")
	}

	isPathExist := false
	for _, route := range routes {
		isPathExist = strings.Contains(route.Path, "/v1")
	}

	assert.Equal(t, isPathExist, true)
}

func TestCreateProductHandler(t *testing.T) {
	testCases := map[string]struct {
		body              string
		returnMockUseCase []any
		expectedCode      int
	}{
		"create product success": {
			body:              `{"id":"123","name":"test","price":100}`,
			returnMockUseCase: []any{&model.Product{Id: "123", Name: "test", Price: 100}, nil},
			expectedCode:      200,
		},
		"create product with invalid body": {
			body:              `true`,
			returnMockUseCase: []any{&model.Product{Id: "123", Name: "test", Price: 100}, nil},
			expectedCode:      400,
		},
		"create product with usecase error": {
			body:              `{"id":"123","name":"test","price":100}`,
			returnMockUseCase: []any{nil, &model.Error{Message: "error"}},
			expectedCode:      500,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			app := gin.Default()
			mockUseCase := mock.NewMockCreateProductUseCase(ctrl)
			applicattion := core.Application{
				CreateProductUseCase: mockUseCase,
			}
			apiController := http.NewApiController(app, applicattion)
			apiController.BindRoutes()

			mockUseCase.
				EXPECT().
				Execute(gomock.Any()).
				Return(tc.returnMockUseCase...).
				AnyTimes()

			req := httptest.NewRequest("POST", "/api/v1/products", strings.NewReader(tc.body))
			resp := httptest.NewRecorder()

			app.ServeHTTP(resp, req)

			assert.Equal(t, resp.Code, tc.expectedCode)
		})
	}
}

func TestGetProductByIdHandler(t *testing.T) {
	testCases := map[string]struct {
		id                string
		returnMockUseCase []any
		expectedCode      int
	}{
		"get product success": {
			id:                "123",
			returnMockUseCase: []any{&model.Product{Id: "123", Name: "test", Price: 100}, nil},
			expectedCode:      200,
		},
		"get product with usecase error": {
			id:                "123",
			returnMockUseCase: []any{nil, &model.Error{Message: "error"}},
			expectedCode:      500,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			app := gin.Default()
			mockUseCase := mock.NewMockGetProductById(ctrl)
			applicattion := core.Application{
				GetProductByIdUseCase: mockUseCase,
			}
			apiController := http.NewApiController(app, applicattion)
			apiController.BindRoutes()

			mockUseCase.
				EXPECT().
				Execute(gomock.Any()).
				Return(tc.returnMockUseCase...).
				AnyTimes()

			req := httptest.NewRequest("GET", "/api/v1/products/"+tc.id, nil)
			resp := httptest.NewRecorder()

			app.ServeHTTP(resp, req)

			assert.Equal(t, resp.Code, tc.expectedCode)
		})
	}
}

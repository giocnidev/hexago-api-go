package usecase_test

import (
	"errors"
	"testing"

	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/giocnidev/api-hexagonal/internal/core/usecase"
	"github.com/giocnidev/api-hexagonal/internal/mock"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestCreateProduct(t *testing.T) {
	testCases := map[string]struct {
		product              model.Product
		returnNockRepository any
		expected             any
	}{
		"create product success": {
			product: model.Product{
				Id:    "123",
				Name:  "test",
				Price: 100,
			},
			returnNockRepository: nil,
			expected:             nil,
		},
		"create product with empty name": {
			product: model.Product{
				Id:    "123",
				Name:  "",
				Price: 100,
			},
			returnNockRepository: nil,
			expected: &model.Error{
				Code:    "",
				Message: "the product name is not correct.",
			},
		},
		"create product with incorrect price": {
			product: model.Product{
				Id:    "123",
				Name:  "test",
				Price: -100,
			},
			returnNockRepository: nil,
			expected: &model.Error{
				Code:    "",
				Message: "the product price is not correct.",
			},
		},
		"create product with repository error": {
			product: model.Product{
				Id:    "123",
				Name:  "test",
				Price: 100,
			},
			returnNockRepository: errors.New("error repository"),
			expected: &model.Error{
				Code:    "",
				Message: "error repository",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock.NewMockProductRepository(ctrl)
			useCase := usecase.NewCreateProductUseCase(mockRepository)

			mockRepository.
				EXPECT().
				Create(gomock.Any()).
				Return(tc.returnNockRepository).
				AnyTimes()

			_, result := useCase.Execute(tc.product)
			assert.Equal(t, result, tc.expected)
		})
	}
}

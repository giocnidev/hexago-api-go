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

func TestGetProductById(t *testing.T) {
	testCases := map[string]struct {
		productId            string
		returnNockRepository []any
		expected             any
	}{
		"get product success": {
			productId:            "123",
			returnNockRepository: []any{&model.Product{}, nil},
			expected:             nil,
		},
		"get product with repository error": {
			productId:            "123",
			returnNockRepository: []any{nil, errors.New("repository error")},
			expected: &model.Error{
				Code:    "",
				Message: "repository error",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock.NewMockProductRepository(ctrl)
			useCase := usecase.NewGetProductByIdUseCase(mockRepository)

			mockRepository.
				EXPECT().
				GetById(gomock.Any()).
				Return(tc.returnNockRepository...).
				AnyTimes()

			_, result := useCase.Execute(tc.productId)
			assert.Equal(t, result, tc.expected)
		})
	}
}

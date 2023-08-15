package temp_test

import (
	"errors"
	"testing"

	"github.com/giocnidev/api-hexagonal/internal/adapter/repository/temp"
	"github.com/giocnidev/api-hexagonal/internal/core/model"
	"github.com/go-playground/assert/v2"
)

func TestCreate(t *testing.T) {
	testCases := map[string]struct {
		product  model.Product
		expected any
	}{
		"create product success": {
			product: model.Product{
				Id:    "123",
				Name:  "test",
				Price: 100,
			},
			expected: nil,
		},
		"create product but already exist": {
			product: model.Product{
				Id:    "001",
				Name:  "test",
				Price: 100,
			},
			expected: errors.New("product already exist"),
		},
	}
	repository := temp.NewTempProductRepository()
	repository.Create(model.Product{
		Id:    "001",
		Name:  "test01",
		Price: 100,
	})
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := repository.Create(tc.product)
			assert.Equal(t, result, tc.expected)
		})
	}
}

func TestGetById(t *testing.T) {
	testCases := map[string]struct {
		productId string
		expected  any
	}{
		"get product success": {
			productId: "001",
			expected:  nil,
		},
		"create product but not exist": {
			productId: "123",
			expected:  errors.New("product not found"),
		},
	}
	repository := temp.NewTempProductRepository()
	repository.Create(model.Product{
		Id:    "001",
		Name:  "test01",
		Price: 100,
	})
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, result := repository.GetById(tc.productId)
			assert.Equal(t, result, tc.expected)
		})
	}
}

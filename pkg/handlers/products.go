package handlers

import (
	"context"

	api_metav1 "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

func CreateProduct(ctx context.Context, product *api.Product) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func ListProducts(ctx context.Context, options *api_metav1.ListOptions) (*api.ProductsList, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func UpdateProduct(ctx context.Context, name string, product *api.Product) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func DeleteProduct(ctx context.Context, name string) error {
	return errors.ErrorNotImplemented.Error()
}

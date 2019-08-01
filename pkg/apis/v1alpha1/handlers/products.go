package handlers

import (
	"context"

	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func ListProducts(ctx context.Context, options *v1.ListOptions) (*api.ProductsList, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

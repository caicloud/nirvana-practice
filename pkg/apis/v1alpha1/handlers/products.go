package handlers

import (
	"context"
	"fmt"

	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

var defaultProduct = &api.Product{
	Metadata: v1.Metadata{
		Name: "Apple",
	},
	Spec: &api.ProductSpec{
		Category: "Fruit",
	},
}

func CreateProduct(ctx context.Context, product *api.Product) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func ListProducts(ctx context.Context, options *v1.ListOptions) (*api.ProductsList, error) {
	results := &api.ProductsList{
		Total: 1,
		Items: []*api.Product{defaultProduct},
	}

	return results, nil
}

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	fmt.Printf("Get product: %s\n", name)

	return defaultProduct, nil
}

func UpdateProduct(ctx context.Context, name string, product *api.Product) (*api.Product, error) {
	fmt.Printf("Update product: %s\n", name)

	return product, nil
}

func DeleteProduct(ctx context.Context, name string) error {
	fmt.Printf("Delete product: %s\n", name)

	return nil
}

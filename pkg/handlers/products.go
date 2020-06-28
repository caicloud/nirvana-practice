package handlers

import (
	"context"

	api_metav1 "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
	product2 "github.com/caicloud/nirvana-practice/pkg/repository/product"
)

type ProductHandler struct {
	repo product2.Repository
}

func NewProductHandler(repo product2.Repository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (p *ProductHandler) Create(ctx context.Context, product *api.Product) (*api.Product, error) {
	product, err := p.repo.Add(product)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (p *ProductHandler) List(ctx context.Context, options *api_metav1.ListOptions) (*api.ProductsList, error) {
	productsSource := p.repo.GetAll(options.Start, options.Limit, options.SortBy, options.ReverseOrder)
	data := &api.ProductsList{
		Total: len(productsSource),
		Items: productsSource,
	}
	return data, nil
}

func (p *ProductHandler) Get(ctx context.Context, name string) (*api.Product, error) {
	product := p.repo.Get(name)
	if product == nil {
		return nil, errors.ErrorNotFound.Error(name)
	}
	return product, nil
}

func (p *ProductHandler) Update(ctx context.Context, name string, product *api.Product) (*api.Product, error) {
	if err := p.repo.Update(name, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductHandler) Delete(ctx context.Context, name string) error {
	if err := p.repo.Delete(name); err != nil {
		return errors.ErrorNotFound.Error(name)
	}
	return nil
}

package handlers

import (
	"context"
	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
	"github.com/caicloud/nirvana-practice/pkg/repository"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (p *ProductHandler) CreateProduct(ctx context.Context, product *api.Product) (*api.Product, error) {
	return p.repo.Add(product), nil
}

func (p *ProductHandler) ListProducts(ctx context.Context, options *v1.ListOptions) (*api.ProductsList, error) {
	productsSource := p.repo.GetAll(options.Start, options.Limit, options.SortBy, options.ReverseOrder)
	data := &api.ProductsList{
		Total: len(productsSource),
		Items: productsSource,
	}
	return data, nil
}

func (p *ProductHandler) GetProduct(ctx context.Context, name string) (*api.Product, error) {
	product := p.repo.Get(name)
	if product == nil {
		return nil, errors.ErrorNotFound.Error(name)
	}
	return product, nil
}

func (p *ProductHandler) UpdateProduct(ctx context.Context, name string, product *api.Product) (*api.Product, error) {
	if err := p.repo.Update(name, product); err != nil {
		return nil, errors.ErrorNotFound.Error(name)
	}
	return product, nil
}

func (p *ProductHandler) DeleteProduct(ctx context.Context, name string) error {
	if err := p.repo.Delete(name); err != nil {
		return errors.ErrorNotFound.Error(name)
	}
	return nil
}

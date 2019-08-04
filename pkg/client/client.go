package client

import (
	"context"
	"fmt"
	"github.com/caicloud/nirvana/definition"
	"net/http"

	"github.com/caicloud/nirvana/rest"

	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

// CacheClient client of cache server
type CacheClient struct {
	*rest.Client
}

// QueryRow returns a product by name
// this can only return one record, when wanting to returns multi-record, to use Query
func (cc *CacheClient) QueryRow(name string) (*v1alpha1.Product, error) {
	if name == "" {
		return nil, errors.ErrorInvalidParameter.Error("name")
	}

	var (
		product *v1alpha1.Product = new(v1alpha1.Product)
	)
	req := cc.Client.Request(http.MethodGet, http.StatusOK, fmt.Sprintf("/api/cache/products/%s", name))
	ctx := context.Background()

	if err := req.Data(product).Do(ctx); err != nil {
		return nil, err
	}

	return product, nil
}

// Query returns multi-record, quantity is ${limit}
func (cc *CacheClient) Query (limit int) ([]*v1alpha1.Product, error) {
	if limit < 0 {
		return nil, errors.ErrorInvalidParameter.Error("limit")
	}

	var (
		productList []*v1alpha1.Product
	)

	if err := cc.Client.Request(http.MethodGet, http.StatusOK, "/api/cache/products").Query("limit", limit).Data(&productList).Do(context.Background()); err != nil {
		return nil, err
	}
	return productList, nil
}

// Create create one record
func (cc *CacheClient) Create(product *v1alpha1.Product) error {
	if product == nil {
		return errors.ErrorInvalidParameter.Error("product")
	}

	return cc.Client.Request(http.MethodPost, http.StatusCreated, "/api/cache/products/"+product.Name).Body(definition.MIMEJSON, product).Do(context.Background())
}

// Update update one record
func (cc *CacheClient) Update(name string, product *v1alpha1.Product) error {
	if name == "" || product == nil {
		return errors.ErrorInvalidParameter.Error("name or product")
	}

	return cc.Request(http.MethodPut, http.StatusOK, "/api/cache/products/" + name).Body(definition.MIMEJSON, product).Do(context.Background())
}

// Delete one record
func (cc *CacheClient) Delete (name string) error {
	if name == "" {
		return errors.ErrorInvalidParameter.Error("name")
	}

	return cc.Request(http.MethodDelete, http.StatusNoContent, "/api/cache/products/"+name).Do(context.Background())
}


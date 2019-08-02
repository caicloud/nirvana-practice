package client

import (
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana/rest"
)

type CacheClient struct {
	*rest.Client
}

func (cc *CacheClient) QueryRow(name string) (*v1alpha1.Product, error) {
	return nil, nil
}

func (cc *CacheClient) Query (limit int) ([]*v1alpha1.Product, error) {
	return nil, nil
}

func (cc *CacheClient) Create(product *v1alpha1.Product) error {
	return nil
}

func (cc *CacheClient) Update(name string, product *v1alpha1.Product) error {
	return nil
}

func (cc *CacheClient) Delete (name string) error {
	return nil
}


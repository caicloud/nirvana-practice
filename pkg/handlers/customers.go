package handlers

import (
	"context"

	v1 "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

func CreateCustomer(ctx context.Context, customer *api.Customer) (*api.Customer, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func ListCustomers(ctx context.Context, options *v1.ListOptions) (*api.CustomersList, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func GetCustomer(ctx context.Context, name string) (*api.Customer, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func UpdateCustomer(ctx context.Context, name string, customer *api.Customer) (*api.Customer, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func DeleteCustomer(ctx context.Context, name string) error {
	return errors.ErrorNotImplemented.Error()
}

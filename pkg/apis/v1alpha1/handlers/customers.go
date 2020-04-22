package handlers

import (
	"context"
	"fmt"

	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

var defaultCustomer = &api.Customer{
	Metadata: v1.Metadata{
		Name: "Tony",
	},
	Spec: &api.CustomerSpec{
		Sex: "Male",
	},
}

func CreateCustomer(ctx context.Context, customer *api.Customer) (*api.Customer, error) {
	return nil, errors.ErrorNotImplemented.Error()
}

func ListCustomers(ctx context.Context, options *v1.ListOptions) (*api.CustomersList, error) {
	results := &api.CustomersList{
		Total: 1,
		Items: []*api.Customer{defaultCustomer},
	}

	return results, nil
}

func GetCustomer(ctx context.Context, name string) (*api.Customer, error) {
	fmt.Printf("Get customer: %s\n", name)

	return defaultCustomer, nil
}

func UpdateCustomer(ctx context.Context, name string, customer *api.Customer) (*api.Customer, error) {
	fmt.Printf("Update customer: %s\n", name)

	return customer, nil
}

func DeleteCustomer(ctx context.Context, name string) error {
	fmt.Printf("Delete customer: %s\n", name)

	return nil
}

package handlers

import (
	"context"

	"github.com/caicloud/nirvana/log"

	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/apis/middlewares"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
)

// GetProduct returns a product named ${name}
func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	client := middlewares.GetCacheClient(ctx)
	if name == "" {
		return nil, errors.ErrorInvalidParameter.Error("name")
	}
	prod, err := client.QueryRow(name)
	if err != nil {
		log.Errorln(err)
		if errors.IsNirvanaError(err) {
			return nil, err
		}
		return nil, errors.ErrorInternal.Error(err.Error())
	}
	return prod, nil
}

// CreateProduct creat a new product
func CreateProduct(ctx context.Context, product *api.Product) error {
	if product == nil {
		return errors.ErrorInvalidParameter.Error()
	}
	err := middlewares.GetCacheClient(ctx).Create(product)
	if err != nil {
		log.Infoln(err)
		if errors.IsNirvanaError(err) {
			return err
		}
		return errors.ErrorInternal.Error(err.Error())
	}
	return nil
}

// UpdateProduct will update the whole product by ${product}
// whose name = ${oldName}
func UpdateProduct(ctx context.Context, oldName string,product *api.Product) error {
	if oldName == "" || product == nil {
		return errors.ErrorInvalidParameter.Error("oldName or product")
	}

	err := middlewares.GetCacheClient(ctx).Update(oldName, product)
	if err != nil {
		log.Infoln(err)
		return errors.ErrorInternal.Error(err.Error())
	}
	return nil
}

// DeleteProduct delete product named ${name}
func DeleteProduct(ctx context.Context, name string) error {
	if name == "" {
		return errors.ErrorInvalidParameter.Error("name")
	}

	err := middlewares.GetCacheClient(ctx).Delete(name)
	if err != nil {
		log.Infoln(err)
		if errors.IsNirvanaError(err) {
			return err
		}
		return errors.ErrorInternal.Error(err.Error())
	}

	return nil
}

// ListProducts returns products list and it's count is limited by options' limit field
func ListProducts(ctx context.Context, options *v1.ListOptions) (*api.ProductsList, error) {
	if options == nil || options.Limit < 0 {
		return nil, errors.ErrorInvalidParameter.Error("options Or options.Limit")
	}

	list, err := middlewares.GetCacheClient(ctx).Query(options)
	if err != nil {
		log.Infoln(err)
		if errors.IsNirvanaError(err) {
			return nil, err
		}
		return nil, errors.ErrorInternal.Error(err.Error())
	}

	pl := &api.ProductsList{
		Items:list,
		Total:len(list),
	}
	return pl, nil
}

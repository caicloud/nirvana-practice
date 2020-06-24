package repository

import (
	"github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/errors"
	"github.com/caicloud/nirvana-practice/pkg/tools/sortKeys"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

type ProductRepository interface {
	Add(product *api.Product) *api.Product
	GetAll(start, limit int, orderKey string, reverseOrder bool) []*api.Product
	Get(name string) *api.Product
	Delete(name string) error
	Update(name string, product *api.Product) error
}

type ProductService struct {
	products []*api.Product
}

func NewProductService() *ProductService {
	products := make([]*api.Product, 0, 10)
	for i := 0; i < 10; i++ {
		labels := make(map[string]string)
		labels["label1"] = "labelTest"
		annotations := make(map[string]string)
		annotations["annotations"] = "annotationsTest"
		creationTimestamp := time.Now()
		soldTimestamp := time.Now().Add(time.Hour * 24)
		var price float64
		price = 22.3
		var sold bool
		sold = true
		product := &api.Product{
			Metadata: v1.Metadata{
				UID:               uuid.NewV4().String(),
				Name:              "product" + strconv.Itoa(i),
				Alias:             "产品" + strconv.Itoa(i),
				Labels:            labels,
				Annotations:       annotations,
				Description:       "这是一个普通的产品",
				CreationTimestamp: &creationTimestamp,
			},
			Spec: &api.ProductSpec{
				Category: "test",
				Price:    &price,
			},
			Status: &api.ProductStatus{
				Sold:          &sold,
				SoldTimestamp: &soldTimestamp,
			},
		}
		products = append(products, product)
	}
	return &ProductService{products: products}
}

func (p *ProductService) Add(product *api.Product) *api.Product {
	uid := uuid.NewV4().String()
	product.UID = uid
	p.products = append(p.products, product)
	return product
}

func (p *ProductService) GetAll(start, limit int, orderKey string, reverseOrder bool) []*api.Product {
	productsBefore := make([]*api.Product, len(p.products))
	copy(productsBefore, p.products)
	sortKeys.SortByKey(productsBefore, orderKey, reverseOrder)
	if start >= len(productsBefore) {
		start = len(productsBefore) - 1
	}
	end := start + limit
	if end > len(productsBefore) {
		end = len(productsBefore)
	}
	return productsBefore[start:end]
}

func (p *ProductService) Get(name string) *api.Product {
	var product *api.Product
	for _, item := range p.products {
		if item.Name == name {
			product = item
			break
		}
	}
	return product
}

func (p *ProductService) Delete(name string) error {
	i := -1
	for index, item := range p.products {
		if item.Name == name {
			i = index
			break
		}
	}
	if i != -1 {
		p.products = append(p.products[:i], p.products[i+1:]...)
		return nil
	}
	return errors.ErrorNotFound.Error(name)
}

func (p *ProductService) Update(name string, product *api.Product) error {
	var update bool
	for i, item := range p.products {
		if item.Name == name {
			product.UID = item.UID
			p.products[i] = product
			update = true
			break
		}
	}
	if update {
		return nil
	}
	return errors.ErrorNotFound.Error(name)

}

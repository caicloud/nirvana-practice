package descriptors

import (
	"github.com/caicloud/nirvana-practice/pkg/apis/cache/handlers"
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"
)

// CacheDescriptor builds and returns a Descriptor for all Product APIs.
func CacheDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path: "/products",
		Definitions: []definition.Definition{
			{
				Method:   definition.List,
				Function: handlers.ListProducts,
				Parameters: []definition.Parameter{
					{
						Source:      definition.Query,
						Name:        "limit",
						Description: "generic list limit",
						Operators:   []definition.Operator{
							validator.Int(""),
						},
					},
				},
				Results:     definition.DataErrorResults("listed products"),
				Description: "list products",
			},
			// TODO
		},
		Children: []definition.Descriptor{
			{
				Path: "/{product}",
				Definitions: []definition.Definition{
					{
						Method: definition.Get,
						Function: handlers.GetProduct,
						Parameters: []definition.Parameter {
							definition.QueryParameterFor("name", "get the production by the name"),
						},
						Results: definition.DataErrorResults("detail of the product"),
						Description: "get one record of production",
					},
					{
						Method: definition.Create,
						Function: handlers.CreateProduct,
						Parameters: []definition.Parameter {
							{
								Source: definition.Body,
								Operators: []definition.Operator {
									validator.Struct(&v1alpha1.Product{}),
								},
							},
						},
						Results: []definition.Result{
							definition.ErrorResult(),
						},
						Description: "create new product",
					},
					{
						Method: definition.Update,
						Function: handlers.UpdateProduct,
						Parameters: []definition.Parameter {
							{
								Source: definition.Body,
								Operators: []definition.Operator {
									validator.Struct(&v1alpha1.Product{}),
								},
								Description: "the update data of product",
							},
							{
								Source: definition.Path,
								Name: "product",
								Operators: []definition.Operator {
									validator.String(""),
								},
								Description: "oldName of the product, because the update operation may change the name",
							},
						},
						Results: []definition.Result{
							definition.ErrorResult(),
						},
						Description: "create new product",
					},
					{
						Method: definition.Delete,
						Function: handlers.DeleteProduct,
						Parameters: []definition.Parameter {
							definition.QueryParameterFor("product", "the name of deleted product"),
						},
						Results: []definition.Result {
							definition.ErrorResult(),
						},
						Description: "delete product",
					},
				},
				Description: "single-target Product APIs",
			},
		},
		Description: "all Product APIs",
	}
}

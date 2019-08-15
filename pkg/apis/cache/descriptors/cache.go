package descriptors

import (
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"

	"github.com/caicloud/nirvana-practice/pkg/apis/cache/handlers"
	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
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
						Source:      definition.Auto,
						Name:        "options",
						Description: "generic list limit",
						Operators:   []definition.Operator{
							validator.Struct(&meta.ListOptions{}),
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
							definition.PathParameterFor("product", "name of product"),
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
							definition.PathParameterFor("product", "oldName of the product, because the update operation may change the name"),
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
							definition.PathParameterFor("product", "name of product"),
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

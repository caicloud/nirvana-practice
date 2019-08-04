package descriptors

import (
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"

	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1/handlers"
)

// ProductDescriptor builds and returns a Descriptor for all Product APIs.
func ProductDescriptor() definition.Descriptor {
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
						Description: "generic list options",
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
						Method:   definition.Get,
						Function: handlers.GetProduct,
						Parameters: []definition.Parameter{
							definition.PathParameterFor("product", "name of the product to get"),
						},
						Results:     definition.DataErrorResults("the get result (or error)"),
						Description: "get product",
					},
					{
						Method: definition.Create,
						Function: handlers.CreateProduct,
						Parameters: []definition.Parameter {
							definition.Parameter{
								Source:      definition.Body,
								Name:        "product",
								Operators:   []definition.Operator{
									validator.Struct(&v1alpha1.Product{}),
								},
								Description: "product of new one",
							},
						},
						Results: []definition.Result {
							definition.ErrorResult(),
						},
						Description: "create production",
					},
					{
						Method: definition.Update,
						Function: handlers.UpdateProduct,
						Parameters: []definition.Parameter {
							definition.Parameter{
								Source:      definition.Path,
								Name:        "product",
								Operators:   []definition.Operator{
									validator.String(""),
								},
								Description: "name of being updated production",
							},
							definition.Parameter{
								Source:      definition.Body,
								Operators:   []definition.Operator{
									validator.Struct(&v1alpha1.Product{}),
								},
								Description: "the whole data of the production that would be updated ",
							},
						},
						Results: []definition.Result {
							definition.ErrorResult(),
						},
						Description: "update production",
					},
					{
						Method: definition.Delete,
						Function: handlers.DeleteProduct,
						Parameters: []definition.Parameter {
							definition.Parameter{
								Source:      definition.Path,
								Name:        "product",
								Operators:   []definition.Operator{
									validator.String(""),
								},
								Description: "name of being deleted product",
							},
						},
						Results: []definition.Result {
							definition.ErrorResult(),
						},
						Description: "delete production",
					},
				},
				Description: "single-target Product APIs",
			},
		},
		Description: "all Product APIs",
	}
}

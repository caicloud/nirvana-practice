package descriptors

import (
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/definition/method"
	"github.com/caicloud/nirvana/operators/validator"

	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/handlers"
	product2 "github.com/caicloud/nirvana-practice/pkg/repository/product"
)

func init() {
	productService := product2.NewService()
	productHandler = handlers.NewProductHandler(productService)
	defaultContainer.Put(productHandler)
	register(products...)
}

var (
	defaultContainer = method.NewContainer()
	productHandler   *handlers.ProductHandler
)

var products = []definition.Descriptor{
	{
		Path:        "/products",
		Description: "Product API",
		Tags:        []string{"product"},
		Definitions: []definition.Definition{listProducts, createProduct},
		Children: []definition.Descriptor{
			{
				Path:        "/{product}",
				Definitions: []definition.Definition{getProduct, updateProduct, deleteProduct},
			},
		},
	},
}

var listProducts = definition.Definition{
	Method:      definition.List,
	Description: "list products",
	Function:    defaultContainer.Get(productHandler, "List"),
	Parameters: []definition.Parameter{
		{
			Source:      definition.Auto,
			Name:        "options",
			Description: "generic list options",
			Operators:   []definition.Operator{validator.Struct(&meta.ListOptions{})},
		},
	},
	Results: definition.DataErrorResults("listed products"),
}

var createProduct = definition.Definition{
	Method:      definition.Create,
	Description: "create product",
	Function:    defaultContainer.Get(productHandler, "Create"),
	Parameters: []definition.Parameter{
		definition.BodyParameterFor("JSON body to describe the new product"),
	},
	Results: definition.DataErrorResults("product"),
}

var getProduct = definition.Definition{
	Method:      definition.Get,
	Description: "get product",
	Function:    defaultContainer.Get(productHandler, "Get"),
	Parameters: []definition.Parameter{
		definition.PathParameterFor("product", "name of the product to get"),
	},
	Results: definition.DataErrorResults("product"),
}

var updateProduct = definition.Definition{
	Method:      definition.Update,
	Description: "update product",
	Function:    defaultContainer.Get(productHandler, "Update"),
	Parameters: []definition.Parameter{
		definition.PathParameterFor("product", "name of the product to update"),
		definition.BodyParameterFor("JSON body to describe the new product"),
	},
	Results: definition.DataErrorResults("product"),
}

var deleteProduct = definition.Definition{
	Method:      definition.Delete,
	Description: "delete product",
	Function:    defaultContainer.Get(productHandler, "Delete"),
	Parameters: []definition.Parameter{
		definition.PathParameterFor("product", "name of the product to delete"),
	},
	Results: []definition.Result{definition.ErrorResult()},
}

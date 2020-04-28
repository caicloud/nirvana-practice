package descriptors

import (
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"

	meta "github.com/caicloud/nirvana-practice/pkg/apis/meta/v1"
	"github.com/caicloud/nirvana-practice/pkg/handlers"
)

func init() {
	register(customers...)
}

var customers = []definition.Descriptor{
	{
		Path:        "/customers",
		Description: "Customer API",
		Tags:        []string{"customer"},
		Definitions: []definition.Definition{listCustomers, createCustomer},
		Children: []definition.Descriptor{
			{
				Path:        "/{customer}",
				Definitions: []definition.Definition{getCustomer, updateCustomer, deleteCustomer},
			},
		},
	},
}

var listCustomers = definition.Definition{
	Method:      definition.List,
	Description: "list customers",
	Function:    handlers.ListCustomers,
	Parameters: []definition.Parameter{
		{
			Source:      definition.Auto,
			Name:        "options",
			Description: "generic list options",
			Operators:   []definition.Operator{validator.Struct(&meta.ListOptions{})},
		},
	},
	Results: definition.DataErrorResults("listed customers"),
}

var createCustomer = definition.Definition{
	Method:      definition.Create,
	Description: "create customer",
	Function:    handlers.CreateCustomer,
	Parameters: []definition.Parameter{
		definition.BodyParameterFor("JSON body to describe the new customer"),
	},
	Results: definition.DataErrorResults("customer"),
}

var getCustomer = definition.Definition{
	Method:      definition.Get,
	Description: "get customer",
	Function:    handlers.GetCustomer,
	Parameters: []definition.Parameter{
		definition.PathParameterFor("customer", "name of the customer to get"),
	},
	Results: definition.DataErrorResults("customer"),
}

var updateCustomer = definition.Definition{
	Method:      definition.Update,
	Description: "update customer",
	Function:    handlers.UpdateCustomer,
	Parameters: []definition.Parameter{
		definition.PathParameterFor("customer", "name of the customer to update"),
		definition.BodyParameterFor("JSON body to describe the new customer"),
	},
	Results: definition.DataErrorResults("customer"),
}

var deleteCustomer = definition.Definition{
	Method:      definition.Delete,
	Description: "delete customer",
	Function:    handlers.DeleteCustomer,
	Parameters: []definition.Parameter{
		definition.PathParameterFor("customer", "name of the customer to delete"),
	},
	Results: []definition.Result{definition.ErrorResult()},
}

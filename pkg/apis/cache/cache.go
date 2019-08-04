package cache

import (
	"sync"

	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
)



// Memory used to store products
// it's key is the product's name
type Cache struct {
	Products map[string]*api.Product
	sync.Mutex
}


package sortKeys

import (
	api "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"sort"
)

type By func(p1, p2 *api.Product) bool

type productSorter struct {
	products []*api.Product
	by       func(p1, p2 *api.Product) bool
}

func (p *productSorter) Len() int {
	return len(p.products)
}

func (p *productSorter) Less(i, j int) bool {
	return p.by(p.products[i], p.products[j])
}

func (p *productSorter) Swap(i, j int) {
	p.products[i], p.products[j] = p.products[j], p.products[i]
}

func (by By) Sort(products []*api.Product) {
	ps := &productSorter{
		products: products,
		by:       by,
	}
	sort.Sort(ps)
}

func SortByKey(products []*api.Product, key string, reverseOrder bool) {
	name := func(p1, p2 *api.Product) bool {
		if reverseOrder {
			return p1.Name > p2.Name
		}
		return p1.Name < p2.Name
	}
	alias := func(p1, p2 *api.Product) bool {
		if reverseOrder {
			return p1.Alias > p2.Alias
		}
		return p1.Alias < p2.Alias
	}
	switch key {
	case "name":
		By(name).Sort(products)
	case "alias":
		By(alias).Sort(products)
	default:
		By(name).Sort(products)
	}

}

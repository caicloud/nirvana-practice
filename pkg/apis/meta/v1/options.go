package v1

// ListOptions includes some generic listing options that can be adopted by most APIs.
type ListOptions struct {
	Start        int    `source:"Query,start,default=0" validate:"gte=0"`
	Limit        int    `source:"Query,limit,default=100" validate:"gt=0"`
	SortBy       string `source:"Query,sortBy"`
	ReverseOrder bool   `source:"Query,reverseOrder"`
}

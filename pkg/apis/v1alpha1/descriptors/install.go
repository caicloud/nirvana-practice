package descriptors

import (
	"github.com/caicloud/nirvana/definition"
)

var descriptors []definition.Descriptor

func register(ds ...definition.Descriptor) {
	descriptors = append(descriptors, ds...)
}

// Descriptor returns a combined descriptor for current version.
func Descriptor() definition.Descriptor {
	return definition.Descriptor{
		Description: "v1alpha1 API",
		Path:        "/v1alpha1",
		Consumes:    []string{definition.MIMEJSON},
		Produces:    []string{definition.MIMEJSON},
		Children:    descriptors,
	}
}

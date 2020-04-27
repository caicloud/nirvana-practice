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
		// Add other content types exclude JSON for consume and produce if necessary,
		// like text, octet-stream etc.
		// Consumes: []string{definition.MIMEOctetStream},
		// Produces: []string{definition.MIMEText},
		Children: descriptors,
	}
}

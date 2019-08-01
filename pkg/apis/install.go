package apis

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/plugins/reqlog"

	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1/descriptors"
)

// Install configures the given Nirvana Config object with the API Descriptor.
func Install(config *nirvana.Config) {
	config.Configure(
		reqlog.Default(),
		nirvana.Descriptor(
			APIDescriptor(),
		),
	)
}

// APIDescriptor builds and returns a descriptor of all APIs.
func APIDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path:     "/api",
		Consumes: []string{definition.MIMEJSON},
		Produces: []string{definition.MIMEJSON},
		Children: []definition.Descriptor{
			{
				Path: "/v1alpha1",
				Children: []definition.Descriptor{
					descriptors.ProductDescriptor(),
				},
				Description: "all v1alpha1 APIs",
			},
		},
		Description: "all APIs",
	}
}

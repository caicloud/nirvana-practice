// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	v1alpha1 "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1/descriptors"

	middleware "github.com/caicloud/go-common/nirvana/middleware"
	def "github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/log"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/api",
		Middlewares: []def.Middleware{
			middleware.Reqlog(log.DefaultLogger()),
		},
		Consumes: []string{def.MIMEJSON},
		Produces: []string{def.MIMEJSON},
		Children: []def.Descriptor{
			v1alpha1.Descriptor(),
		},
	}
}

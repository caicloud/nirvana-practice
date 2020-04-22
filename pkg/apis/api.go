// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	v1alpha1 "github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1/descriptors"

	def "github.com/caicloud/nirvana/definition"
)

// Descriptor returns a combined descriptor for APIs of all versions.
func Descriptor() def.Descriptor {
	return def.Descriptor{
		Description: "APIs",
		Path:        "/api",
		Consumes:    []string{def.MIMEJSON},
		Produces:    []string{def.MIMEJSON},
		Children: []def.Descriptor{
			v1alpha1.Descriptor(),
		},
	}
}

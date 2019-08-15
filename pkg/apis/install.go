package apis

import (
	"context"
	"sync"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/plugins/reqlog"
	"github.com/caicloud/nirvana/rest"

	"github.com/caicloud/nirvana-practice/pkg/apis/cache"
	descriptors2 "github.com/caicloud/nirvana-practice/pkg/apis/cache/descriptors"
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/caicloud/nirvana-practice/pkg/apis/v1alpha1/descriptors"
	client2 "github.com/caicloud/nirvana-practice/pkg/client"
)

// Install configures the given Nirvana Config object with the API Descriptor.
func Install(config *nirvana.Config, cacheConfig *rest.Config) {
	config.Configure(
		reqlog.Default(),
		nirvana.Descriptor(
			definition.Descriptor{
				Path:        "/",
				Middlewares: []definition.Middleware{
					func(ctx context.Context, chain definition.Chain) error {
						client, err := rest.NewClient(cacheConfig)
						if err != nil {
							return err
						}
						cli := client2.NewCacheClient(client)
						return chain.Continue(context.WithValue(ctx, "client", cli))
					},
				},
				Children:    []definition.Descriptor{
					APIDescriptor(),
				},
				Description: "install client in context",
			},
		),
	)
}

// Install configures the given Nirvana Config object with the API Descriptor.
func InstallCache(config *nirvana.Config) {
	config.Configure(
		reqlog.Default(),
		nirvana.Descriptor(
			CacheAPIDescriptor(),
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

// CacheAPIDescriptor builds and returns a descriptor of all Cache APIs.
func CacheAPIDescriptor() definition.Descriptor {
	c := cache.Cache{
		Products: map[string]*v1alpha1.Product{},
		Mutex: sync.Mutex{},
	}

	return definition.Descriptor{
		Path:     "/api",
		Consumes: []string{definition.MIMEJSON},
		Produces: []string{definition.MIMEJSON},
		Middlewares: []definition.Middleware {
			// register cache
			func(ctx context.Context, chain definition.Chain) error {
				return chain.Continue(context.WithValue(ctx, "cache", &c))
			},
		},
		Children: []definition.Descriptor{
			{
				Path: "/cache",
				Children: []definition.Descriptor{
					descriptors2.CacheDescriptor(),
				},
				Description: "all cache APIs",
			},
		},
		Description: "all cache APIs",
	}
}
package middlewares

import (
	"context"

	"github.com/caicloud/nirvana-practice/pkg/apis/cache"
	"github.com/caicloud/nirvana-practice/pkg/client"
)

// GetCache returns the cache instance
// this function should use in cache server
func GetCache(ctx context.Context) *cache.Cache {
	return ctx.Value("cache").(*cache.Cache)
}

// GetCacheClient returns cache client
func GetCacheClient(ctx context.Context) *client.CacheClient {
	return ctx.Value("client").(*client.CacheClient)
}
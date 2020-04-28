package middleware

import (
	"context"
	"time"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/service"
)

// Reqlog middleware functions similarly to the reqlog plugin from Nirvana. It prints the basic
// information on every request and response that passes through it.
func Reqlog(logger log.Logger) definition.Middleware {
	return func(ctx context.Context, chain definition.Chain) error {
		start := time.Now()
		httpCtx := service.HTTPContextFrom(ctx)

		err := chain.Continue(ctx)

		request := httpCtx.Request()
		response := httpCtx.ResponseWriter()
		logger.Infoln(
			request.Method,
			response.StatusCode(),
			response.ContentLength(),
			time.Since(start).String(),
			request.URL.String(),
		)

		return err
	}
}

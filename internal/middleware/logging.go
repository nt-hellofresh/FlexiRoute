package middleware

import (
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"log"
)

func LoggingMiddleware(handler flexiroute.ApiHandler) flexiroute.ApiHandler {
	return func(ctx *flexiroute.Context) error {
		log.Printf(
			"- [%v] %v",
			ctx.RequestMethod(),
			ctx.RequestURL().Path,
		)
		return handler(ctx)
	}
}

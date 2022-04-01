package cfrouter

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

func (cfr *CfRouter) Use(mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		cfr.middlewares = append(cfr.middlewares, fn)
	}
}

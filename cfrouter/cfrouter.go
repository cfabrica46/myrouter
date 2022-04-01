package cfrouter

import (
	"log"
	"net/http"
)

type CfRouter struct {
	routes      []*Route
	middlewares []middleware
}

func NewCfRouter() *CfRouter {
	return &CfRouter{}
}

func (cfr *CfRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler

	path := r.URL.Path
	if path == "" {
		path = "/"
	}

	if path[0] != '/' {
		path = "/" + path
	}

	for _, route := range cfr.routes {
		if route.err != nil {
			log.Panic(route.err)
		}

		if path == route.path && route.methods[r.Method] {
			for i := len(cfr.middlewares) - 1; i >= 0; i-- {
				if i == len(cfr.middlewares)-1 {
					handler = cfr.middlewares[i].Middleware(route.handler)
				} else {
					handler = cfr.middlewares[i].Middleware(handler)
				}
			}
		}
	}

	handler.ServeHTTP(w, r)
}

func (cfr *CfRouter) NewRoute() *Route {
	route := &Route{methods: make(map[string]bool)}

	cfr.routes = append(cfr.routes, route)

	return route
}

func (cfr *CfRouter) Methods(methods ...string) *Route {
	return cfr.NewRoute().Methods(methods...)
}

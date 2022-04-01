package cfrouter

import (
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	methods map[string]bool
	handler http.Handler
	path    string
	err     error
}

func (r *Route) Methods(methods ...string) *Route {
	for k := range methods {
		methods[k] = strings.ToUpper(methods[k])
		r.methods[methods[k]] = true
	}

	return r
}

func (r *Route) SetHandler(handler http.Handler) *Route {
	if r.err == nil {
		r.handler = handler
	}

	return r
}

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route {
	return r.SetHandler(http.HandlerFunc(f))
}

func (r *Route) Path(path string) *Route {
	if r.err != nil {
		return r
	}

	if len(path) > 0 && path[0] != '/' {
		r.err = fmt.Errorf("mux: path must start with a slash, got %q", path)

		return r
	}

	r.path = path

	return r
}

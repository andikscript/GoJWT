package service

import "net/http"

type CustomMux struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middleware = append(c.middleware, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, v := range c.middleware {
		current = v(current)
	}

	current.ServeHTTP(w, r)
}

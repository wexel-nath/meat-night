package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/api/handler"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range getRoutes() {
		router.Handle(route.method, route.pattern, middleware(route.handler))
	}

	return router
}

type route struct {
	method  string
	pattern string
	handler httprouter.Handle
}

func getRoutes() []route {
	return []route{
		{
			method:  http.MethodGet,
			pattern: "/healthz",
			handler: handler.HealthzHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/upcoming-hosts",
			handler: handler.UpcomingHostsHandler,
		},
	}
}

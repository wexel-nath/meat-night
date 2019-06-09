package api

import (
	"net/http"

	"github.com/wexel-nath/meat-night/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func healthzHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		logger.Error(err)
	}
}

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
			handler: healthzHandler,
		},
	}
}

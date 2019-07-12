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
			handler: requestHandler(handler.Healthz),
		},
		{
			method:  http.MethodGet,
			pattern: "/mateos",
			handler: requestHandler(handler.ListMateos),
		},
		{
			method:  http.MethodPost,
			pattern: "/dinner",
			handler: requestHandler(handler.CreateDinner),
		},
		{
			method:  http.MethodPost,
			pattern: "/schedule",
			handler: requestHandler(handler.Schedule),
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/accept",
			handler: requestHandler(handler.AcceptHostInvite),
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/decline",
			handler: requestHandler(handler.DeclineHostInvite),
		},
		{
			method:  http.MethodGet,
			pattern: "/guest/:inviteID/accept",
			handler: requestHandler(handler.AcceptGuestInvite),
		},
		{
			method:  http.MethodGet,
			pattern: "/guest/:inviteID/decline",
			handler: requestHandler(handler.DeclineGuestInvite),
		},
	}
}

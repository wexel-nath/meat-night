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
			pattern: "/mateos",
			handler: handler.ListMateosHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/dinner",
			handler: handler.CreateDinnerHandler,
		},
		{
			method:  http.MethodPost,
			pattern: "/schedule",
			handler: handler.ScheduleHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/accept",
			handler: handler.AcceptHostInviteHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/decline",
			handler: handler.DeclineHostInviteHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/guest/:inviteID/accept",
			handler: handler.AcceptGuestInviteHandler,
		},
		{
			method:  http.MethodGet,
			pattern: "/guesst:inviteID/decline",
			handler: handler.DeclineGuestInviteHandler,
		},
	}
}

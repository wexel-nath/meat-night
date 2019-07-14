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
			handler: jsonResponseHandler(handler.Healthz),
		},
		{
			method:  http.MethodGet,
			pattern: "/mateos",
			handler: jsonResponseHandler(handler.ListMateos),
		},
		{
			method:  http.MethodPost,
			pattern: "/dinner",
			handler: jsonResponseHandler(handler.CreateDinner),
		},
		{
			method:  http.MethodPost,
			pattern: "/schedule",
			handler: jsonResponseHandler(handler.Schedule),
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/accept",
			handler: giphyResponseHandler(handler.AcceptHostInvite), // don't serve giphy??
		},
		{
			method:  http.MethodGet,
			pattern: "/host/:inviteID/decline",
			handler: giphyResponseHandler(handler.DeclineHostInvite),
		},
		{
			method:  http.MethodGet,
			pattern: "/guest/:inviteID/accept",
			handler: giphyResponseHandler(handler.AcceptGuestInvite),
		},
		{
			method:  http.MethodGet,
			pattern: "/guest/:inviteID/decline",
			handler: giphyResponseHandler(handler.DeclineGuestInvite),
		},
		{
			method:  http.MethodPost,
			pattern: "/dinner/update",
			handler: giphyResponseHandler(handler.UpdateVenue),
		},
		{
			method:  http.MethodGet,
			pattern: "/dinner/:dinnerID/update",
			handler: handler.UpdateVenueForm,
		},
	}
}

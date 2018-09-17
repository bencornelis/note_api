package app

import (
	"fmt"
	"net/http"

	"github.com/bencornelis/note_api/util"
)

const (
	// http methods
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Route struct {
	Pattern     string
	Method      string
	Handler     func(w http.ResponseWriter, r *http.Request)
	Middlewares []util.Middleware
}

func (app *App) setupRoutes() {
	routes := app.routes()
	baseMiddlewares := []util.Middleware{}
	app.attachRoutes(routes, baseMiddlewares)
}

func (app *App) attachRoutes(routes []Route, baseMiddlewares []util.Middleware) {
	apiPrefix := "/api"
	router := app.router.PathPrefix(apiPrefix).Subrouter()

	fmt.Println("routes:")
	for _, route := range routes {
		var handler http.Handler
		handler = http.HandlerFunc(route.Handler)
		middlewares := append(route.Middlewares, baseMiddlewares...)
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}

		fmt.Println(apiPrefix+route.Pattern, route.Method)

		router.
			Path(route.Pattern).
			Methods(route.Method).
			Handler(handler)
	}
}

func (app *App) routes() []Route {
	return []Route{
		// notes
		Route{
			"/notes",
			GET,
			app.GetNotes,
			[]util.Middleware{util.AuthMiddleware},
		},
		Route{
			"/notes/{id:[0-9]+}",
			GET,
			app.GetNote,
			[]util.Middleware{util.AuthMiddleware},
		},
		Route{
			"/notes",
			POST,
			app.CreateNote,
			[]util.Middleware{util.AuthMiddleware},
		},
		Route{
			"/notes/{id:[0-9]+}",
			PATCH,
			app.UpdateNote,
			[]util.Middleware{util.AuthMiddleware},
		},
		Route{
			"/notes/{id:[0-9]+}",
			DELETE,
			app.DeleteNote,
			[]util.Middleware{util.AuthMiddleware},
		},
		// auth
		Route{
			"/signup",
			POST,
			app.Signup,
			[]util.Middleware{},
		},
		Route{
			"/login",
			POST,
			app.Login,
			[]util.Middleware{},
		},
	}
}

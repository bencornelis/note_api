package app

import (
	"net/http"

	"github.com/bencornelis/note_api/model"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	store  *model.DatabaseStore
}

func NewApp(r *mux.Router, s *model.DatabaseStore) *App {
	app := &App{r, s}
	return app
}

func (app *App) ListenAndServe(addr string) error {
	app.setupRoutes()

	return http.ListenAndServe(addr, app.router)
}

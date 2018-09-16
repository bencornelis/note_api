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
	app.router.HandleFunc("/notes", app.GetNotes).Methods("GET")
	app.router.HandleFunc("/notes/{id:[0-9]+}", app.GetNote).Methods("GET")
	app.router.HandleFunc("/notes", app.CreateNote).Methods("POST")
	app.router.HandleFunc("/notes/{id:[0-9]+}", app.UpdateNote).Methods("PATCH")
	app.router.HandleFunc("/notes/{id:[0-9]+}", app.DeleteNote).Methods("DELETE")

	return http.ListenAndServe(addr, app.router)
}

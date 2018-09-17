package app

import (
	"encoding/json"
	"net/http"

	"github.com/bencornelis/note_api/model"
	"github.com/bencornelis/note_api/util"
)

func (app *App) GetNotes(w http.ResponseWriter, r *http.Request) {
	userId := util.ParseUserId(r)

	notes, err := app.store.GetNotes(userId)
	if err != nil {

	}

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) GetNote(w http.ResponseWriter, r *http.Request) {
	userId := util.ParseUserId(r)

	id, err := util.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note, err := app.store.GetNote(id, userId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) CreateNote(w http.ResponseWriter, r *http.Request) {
	userId := util.ParseUserId(r)

	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note.UserID = userId
	note, err := app.store.CreateNote(note)
	if err != nil {

	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) UpdateNote(w http.ResponseWriter, r *http.Request) {
	userId := util.ParseUserId(r)

	id, err := util.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note.UserID = userId
	note, err = app.store.UpdateNote(id, note)
	if err != nil {

	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) DeleteNote(w http.ResponseWriter, r *http.Request) {
	userId := util.ParseUserId(r)

	id, err := util.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = app.store.DeleteNote(id, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

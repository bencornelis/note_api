package app

import (
	"encoding/json"
	"net/http"

	"github.com/bencornelis/note_api/model"
	"github.com/bencornelis/note_api/utils"
)

func (app *App) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := app.store.GetNotes()
	if err != nil {

	}

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note, err := app.store.GetNote(id)

	if err := json.NewEncoder(w).Encode(note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note, err := app.store.CreateNote(note)
	if err != nil {

	}

	payload := map[string]uint{"noteId": note.ID}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	note, err = app.store.UpdateNote(id, note)
	if err != nil {

	}

	payload := map[string]uint{"noteId": note.ID}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = app.store.DeleteNote(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

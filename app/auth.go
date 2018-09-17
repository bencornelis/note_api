package app

import (
	"encoding/json"
	"net/http"

	"github.com/bencornelis/note_api/model"
	"github.com/bencornelis/note_api/util"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	password := user.Password

	user, err := app.store.FindUser(user.Username)
	if err != nil {

	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {

	}

	token, err := util.GenerateToken(&user)
	if err != nil {

	}

	payload := make(map[string]string)
	payload["token"] = token
	if err := json.NewEncoder(w).Encode(token); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *App) Signup(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	user, err := app.store.CreateUser(user)
	if err != nil {

	}

	token, err := util.GenerateToken(&user)
	if err != nil {

	}

	payload := make(map[string]string)
	payload["token"] = token
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

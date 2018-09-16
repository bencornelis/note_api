package main

import (
	"fmt"
	"log"

	"github.com/bencornelis/note_api/app"
	"github.com/bencornelis/note_api/model"
	"github.com/gorilla/mux"
)

const (
	port = 8000
)

func main() {
	s, err := model.NewDatabaseStore()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	defer s.Close()

	r := mux.NewRouter()
	app := app.NewApp(r, s)

	log.Fatal(app.ListenAndServe(fmt.Sprintf(":%d", port)))
}

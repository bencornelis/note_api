package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseId(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

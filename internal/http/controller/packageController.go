package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// HTTP handler accessing data from the request context.
func GetPackageHandler(w http.ResponseWriter, r *http.Request) {
	// respond to the client
	w.Write([]byte(fmt.Sprintf("hi")))
}

func CreatePackageHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	if body == nil {
		w.Write([]byte(""))
	}
}

func UpdatePackageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	version := chi.URLParam(r, "VERSION")
	w.Write([]byte(fmt.Sprintf("You want to delete package with id %s and version %s", id, version)))
}

func DeletePackageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	version := chi.URLParam(r, "VERSION")
	w.Write([]byte(fmt.Sprintf("You want to delete package with id %s and version %s", id, version)))
}

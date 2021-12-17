package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) router() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/index", app.index)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	return mux
}

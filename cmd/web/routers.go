package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) router() *mux.Router {
	//fetch static files
	fileServer := http.FileServer(http.Dir("C:/Users/P. Zion/go/src/notebin/ui/static"))
	mux := mux.NewRouter()

	mux.HandleFunc("/index", app.index)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	http.Handle("/", mux)
	return mux
}

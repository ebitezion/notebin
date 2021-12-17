package main

import (
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "%s", "Hello World")
	app.template.t.Execute(w, nil)
}

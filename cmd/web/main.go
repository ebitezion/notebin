package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type application struct {
	template td
	infoLog  log.Logger
	errorLog log.Logger
}

type td struct {
	t    *template.Template
	data interface{}
}

var app application
var t *template.Template

func main() {
	//attach template
	t, err := template.ParseFiles("../../ui/html/index.html")
	if err != nil {
		log.Fatal("Template Error", err)
	}
	app.template.t = t

	//run server here
	err = server()
	if err != nil {
		log.Fatal("Server failed :", err)
	}
}

func server() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: app.router(),
	}
	return srv.ListenAndServe()
}

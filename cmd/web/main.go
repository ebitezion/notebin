package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	template td
	infoLog  *log.Logger
	errorLog *log.Logger
	cfg      Config
}

type td struct {
	t    *template.Template
	data *interface{}
}
type Config struct {
	Port      string
	StaticDir string
}

var (
	app   application
	files []string
)

//var t *template.Template
func init() {

}
func main() {
	//
	files = []string{

		"../../ui/html/home.page.tmpl",
		"../../ui/html/base.layout.tmpl",
		"../../ui/html/footer.partial.tmpl",
	}
	//attach template
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("Template Error ", err)

	}
	app.template.t = t

	//flags
	cfg := new(Config)
	flag.StringVar(&cfg.Port, "Port", ":8080", "server port")
	flag.Parse()

	//Initialize the log
	app.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//run server here
	app.infoLog.Println("Notebin server runnng on port", cfg.Port)
	err = app.server()
	if err != nil {
		app.errorLog.Fatal("Server failed: ", err)
	}

}

func (app *application) server() error {
	srv := &http.Server{
		Addr:     ":8080",
		Handler:  app.router(),
		ErrorLog: app.errorLog,
	}
	return srv.ListenAndServe()
}

//todo:create command line flags done
//todo:do templates
//todo:templates cache
//todo:middlewares
//todo:redis
//todo:RDB
//todo:make pages
//todo:Login using JWT
//todo:social login
//todo:Registarion
//todo:context
//todo:test and performance
//todo:microservice
//todo:consume inbuilt func
//todo:publish

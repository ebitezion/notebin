package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	DB        *sql.DB
}

var (
	app application

	files []string
)

/*
 The database constants
*/
const (
	Port   = ":8080"
	DBhost = "localhost"
	DBuser = "root"

	DBport = "3306"
	DBname = "snippetdb"
)

func init() {
	//connect to db

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?parseTime=true", DBuser, DBhost, DBport, DBname)
	db, err := OpenDb(dsn)
	if err != nil {
		log.Fatalln("db  conn error ", err.Error())
		return
	}
	//add to application struct
	app.cfg.DB = db
	//ping db to test connection
	if err := db.Ping(); err != nil {
		fmt.Println("Ping connection error ", err)
	} else {
		fmt.Println("DB pinged")
	}

}
func main() {

	//templates
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
func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//ping db
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
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
//todo:do templates done
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

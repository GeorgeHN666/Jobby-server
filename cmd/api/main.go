package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GeorgeHN666/Jobby-server/internals/routes"
)

const VERSION = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		URI string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           routes.Serve(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	app.infoLog.Printf("Server starting at PORT %d in %s mode", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 3500, "port")
	flag.Parse()

	cfg.db.URI = os.Getenv("DB")
	cfg.env = os.Getenv("ENV")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  VERSION,
	}

	err := app.Serve()
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}

}

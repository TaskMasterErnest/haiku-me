package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//a struct to hold all global dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address") //adding flags for the command line
	flag.Parse()

	//adding custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	//a struct containing the application dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//initializing a struct to house parameters to be served
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), //app.routes() contains the servemux
	}

	//serving the application
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

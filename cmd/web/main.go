package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//a struct to hold all global dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address") //adding flags for the command line
	dsn := flag.String("dsn", "web:Mysql/pass1@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	//adding custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	//initializing a database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close() //close connection before main() exits

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// connecting the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

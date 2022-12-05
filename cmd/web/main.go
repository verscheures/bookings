package main

import (
	"encoding/gob"
	"log"
	"os"

	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/verscheures/bookings/internal/config"
	"github.com/verscheures/bookings/internal/driver"
	"github.com/verscheures/bookings/internal/handlers"
	"github.com/verscheures/bookings/internal/helpers"
	"github.com/verscheures/bookings/internal/models"
	"github.com/verscheures/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application functio
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	log.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}


func run()(*driver.DB, error) {
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog =  log.New(os.Stdout,"ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// What will we store in session
	gob.Register(models.Reservation{})


	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connect to db
	log.Println("Connecting to DB")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=kuiken password=Start123")
	if err != nil {
		log.Fatal("Cannot connect to database! dying...")
	}
	log.Println("Connected to DB")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return nil,err
	}
	app.TemplateCache = tc
	app.UseCache = false
	helpers.NewHelpers(&app)
	// repository setup
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	return db,nil
}
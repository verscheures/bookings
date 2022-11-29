package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/verscheures/bookings/internal/config"
	"github.com/verscheures/bookings/internal/models"
)


var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M){

	testApp.InProduction = false
	// What will we store in session
	gob.Register(models.Reservation{})


	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {}

func (tw *myWriter) Header() http.Header{
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(int){
}

func (tw *myWriter) Write(b []byte)(int, error){
	length := len(b)
	return length, nil
}
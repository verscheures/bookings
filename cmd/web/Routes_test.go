package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/verscheures/bookings/internal/config"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)
switch v := mux.(type) {
case *chi.Mux:
	// all good
default:
	t.Errorf("Type is not *chi.Mux, type is %T", v)
}

}
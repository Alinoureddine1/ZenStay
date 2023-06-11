package main

import (
	"testing"

	"github.com/Alinoureddine1/ZenStay/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	// call the routes function which returns a ServeMux
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf("type is not *http.ServeMux, but is %T", v)
	}

}

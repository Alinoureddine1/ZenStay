package main

import (
	"net/http"

	"github.com/Alinoureddine1/ZenStay/pkg/config"
	"github.com/Alinoureddine1/ZenStay/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	//http handler

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

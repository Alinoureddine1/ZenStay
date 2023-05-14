package main

import (
	"net/http"

	"github.com/Alinoureddine1/ZenStay/pkg/config"
	"github.com/Alinoureddine1/ZenStay/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//http handler

	mux := chi.NewRouter()

	//Recovers instead of panicking
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

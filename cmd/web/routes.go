package main

import (
	"net/http"

	"github.com/Alinoureddine1/ZenStay/pkg/config"
	"github.com/Alinoureddine1/ZenStay/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	//http handler
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}

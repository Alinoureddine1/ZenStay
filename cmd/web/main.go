package main

import (
	"log"
	"net/http"

	"github.com/Alinoureddine1/ZenStay/pkg/config"
	"github.com/Alinoureddine1/ZenStay/pkg/handlers"
	"github.com/Alinoureddine1/ZenStay/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(portNumber, nil)
}

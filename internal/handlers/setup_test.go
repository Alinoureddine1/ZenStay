package handlers

import (
	"encoding/gob"
	"log"
	"net/http"

	"time"

	"github.com/Alinoureddine1/ZenStay/internal/config"
	"github.com/Alinoureddine1/ZenStay/internal/models"
	"github.com/Alinoureddine1/ZenStay/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache(pathToTemplates)
	if err != nil {
		log.Fatal("Cannot create template cache")

	}

	app.TemplateCache = tc
	app.UseCache = true
	repo := NewRepo(&app)
	NewHandlers(repo)
	render.NewTemplates(&app)
	mux := chi.NewRouter()

	//Recovers instead of panicking
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)

	mux.Get("/contact", Repo.Contact)
	mux.Get("/search-availability", Repo.SearchAvailability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	mux.Get("/royal-suite", Repo.RoyalSuites)
	mux.Get("/deluxe-suite", Repo.DeluxeSuites)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// func CreateTestTemplateCache() (map[string]*template.Template, error) {

// 	myCache := map[string]*template.Template{}

// 	pages, err := filepath.Glob(fmt.Sprint("%s/*.page.tmpl", pathToTemplates))
// 	if err != nil {
// 		return myCache, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }

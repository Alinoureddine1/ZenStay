package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Alinoureddine1/ZenStay/internal/config"
	"github.com/justinas/nosurf"

	"github.com/Alinoureddine1/ZenStay/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get the requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Println("Could not get template from cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	//return the template to the browser
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the pages
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Println("Error parsing template:", err)
		return myCache, err
	}
	// loop through the pages
	for _, page := range pages {
		name := filepath.Base(page)
		// parse the page template in the layout
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("Error parsing template:", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("Error parsing template:", err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("Error parsing template:", err)
				return myCache, err
			}
		}
		// add the template to the cache
		myCache[name] = ts
	}
	return myCache, nil
}

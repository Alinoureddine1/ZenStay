package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Alinoureddine1/ZenStay/pkg/config"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

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

	_ = t.Execute(buf, nil)

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

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check if the template is in the cache
// 	_, ok := tc[t]

// 	if !ok {
// 		// if not, create the template cache
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println("Error parsing template:", err)
// 		}

// 	} else {
// 		// get the template from the cache
// 		log.Println("Template from cache")
// 	}
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error parsing template:", err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	// cache the template
// 	tc[t] = tmpl
// 	return nil

// }

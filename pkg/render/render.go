package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if the template is in the cache
	_, ok := tc[t]

	if !ok {
		// if not, create the template cache
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("Error parsing template:", err)
		}

	} else {
		// get the template from the cache
		log.Println("Template from cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// cache the template
	tc[t] = tmpl
	return nil

}

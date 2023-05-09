package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

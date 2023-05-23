package models

import "github.com/Alinoureddine1/ZenStay/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //Data map holds all types of data
	CSRFToken string                 //Cross Site Request Forgery token
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

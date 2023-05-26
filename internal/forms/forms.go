package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	Errors errors
	url.Values
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a custom Form struct
func New(data url.Values) *Form {
	return &Form{
		errors(map[string][]string{}),
		data,
	}
}

func (f *Form) Has(field string, r *http.Request) bool {

	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// MinLength checks for string minimum length
func (f *Form) MinLength(field string, r *http.Request, min int) bool {

	x := r.Form.Get(field)
	if len(x) < min {
		f.Errors.Add(field, "This field must be at least "+string(min)+" characters long")
		return false
	}
	return true
}

package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	Errors errors
	url.Values
}

// New initializes a custom Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
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

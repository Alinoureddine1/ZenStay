package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	Key string
	Val string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, 200},
	{"contact", "/contact", "GET", []postData{}, 200},
	{"royal-suite", "/royal-suite", "GET", []postData{}, 200},
	{"deluxe-suite", "/deluxe-suite", "GET", []postData{}, 200},
	{"make-reservation", "/make-reservation", "GET", []postData{}, 200},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{Key: "first_name", Val: "Tester"},
		{Key: "last_name", Val: "1"},
		{Key: "email", Val: "test@test.com"},
		{Key: "phone", Val: "123456789"},
	}, 200},
	{"post-make-reservation-json", "/make-reservation-json", "POST", []postData{
		{Key: "first_name", Val: "Tester"},
		{Key: "last_name", Val: "1"},
		{Key: "email", Val: "test@test.com"},
		{Key: "phone", Val: "123456789"},
	}, 200},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	// create a new server using the httptest package
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {

		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {

		}

	}

}

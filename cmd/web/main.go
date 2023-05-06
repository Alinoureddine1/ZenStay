package main

import (
	"net/http"

	"github.com/Alinoureddine1/ZenStay/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(portNumber, nil)
}

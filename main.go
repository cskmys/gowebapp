package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)           // for Gorilla mux "/" does not mean route all the URLs, it only means route "/" URL
	router.HandleFunc("/contact", contact) // "/contact" means only "/contact" not "/contact/"
	// for any URL other than "/" and "/contact" mux throws its own default 404 page(even "/contact/" will give 404 page)

	http.ListenAndServe(":3000", router)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		// this part will not run as Gorilla mux will throw its own 404 page in case of no url match
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Can't find the page you are looking for</h1><p>email me if you keep seeing this</p>")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc) // for Gorilla mux "/" does not mean route all the URLs, it only means route "/" URL
	router.HandleFunc("/contact", handleFunc)
	// for any URL other than "/" and "/contact" mux throws its own default 404 page(even "/contact/" will give 404 page)

	http.ListenAndServe(":3000", router)
}

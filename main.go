package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" { // homepage URL "/"
		fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
	} else if r.URL.Path == "/contact" { // contact page URL "/contact"
		fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} // if the URL is incorrect, no page is served
	// now we are serving html based on url accessed in other words we are adding pages
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

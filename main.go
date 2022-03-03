package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!!!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send me an email at <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page...</h1>")
}

func custom404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Cannot find the page you are looking for</h1><p>email me if you keep seeing this</p>")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact) // here "/contact" and "/contact/" get routed to same function "contact"
	router.GET("/faq", faq)
	router.NotFound = http.HandlerFunc(custom404)
	http.ListenAndServe(":3000", router)
}

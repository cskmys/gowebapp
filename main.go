// using "html/template" to prevent html injection attack

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template // just for now, we are using global variables to keep things simple
// but in the later stage where we get code ready for production, we will remove them and clean up
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil { // passing "nil" as there is no dynamic data to insert into the template
		panic(err) // if you don't panic, you may see weird or half-rendered html page which makes you wonder where the bug is
		// hence better to panic here
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
}

func custom404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Can't find the page you are looking for</h1><p>email me if you keep seeing this</p>")
}

type User struct {
	Name string
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml") // just for now, we'll need to execute the code from the root of this directory as relative path is used
	// will change using the relative path in the future
	if err != nil {
		panic(err) // normally you should be handling errors more gracefully but for an unrecoverable error such as this panic will do
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	http.ListenAndServe(":3000", router)
}

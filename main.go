// using "html/template" to prevent html injection attack

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate, contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err) // not a good idea to throw panic when program is running, but it is important to check for error otherwise
		// you might not know why the page is rendered weirdly or partially
		// for now we just go with panic and at a later time we will clean it up with a more graceful error handling code
	}
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

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err) // it is ok to throw panic in initializing code as failure in initialization is normally unrecoverable
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	http.ListenAndServe(":3000", router)
}

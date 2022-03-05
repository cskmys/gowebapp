package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var homeTemplate, contactTemplate, custom404Template *template.Template

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func custom404(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := custom404Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml", "views/layouts/footer.gohtml") // passing "layouts/footer.gohtml" to link and use the template defined in it
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	custom404Template, err = template.ParseFiles("views/custom404.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	log.Fatal(http.ListenAndServe(":3000", router))
}

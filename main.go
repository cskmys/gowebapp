package main

import (
	"github.com/gorilla/mux"
	"gowebapp/views"
	"log"
	"net/http"
)

var homeView, contactView, custom404View *views.View

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func custom404(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := custom404View.Render(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	custom404View = views.NewView("bootstrap", "views/custom404.gohtml")

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	log.Fatal(http.ListenAndServe(":3000", router))
}

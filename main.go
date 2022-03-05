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
	// all gohtml files and templates are compiled during "homeView" creation
	// now doing "ExecuteTemplate" will execute "bootstrap" from "views/layout/bootstrap.gohtml" which will call "yield" template
	// for "homeView", "yield" is present in "views/home.gohtml" which will in turn call "footer" present in "views/layout/footer.gohtml"
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func custom404(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := custom404View.Template.ExecuteTemplate(w, custom404View.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	// "bootstrap" template's file "/views/layouts/bootstrap.gohtml" is hard coded to compile inside "views.NewView" along with other template files
	// hence, just using it by its name rather than the file name is not a problem
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	custom404View = views.NewView("bootstrap", "views/custom404.gohtml")

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	log.Fatal(http.ListenAndServe(":3000", router))
}

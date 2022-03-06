package main

import (
	"github.com/gorilla/mux"
	"gowebapp/controllers"
	"gowebapp/views"
	"net/http"
)

var homeView, contactView, custom404View *views.View

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func custom404(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(custom404View.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	custom404View = views.NewView("bootstrap", "views/custom404.gohtml")

	usersC := controllers.NewUsers()

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", usersC.New)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	must(http.ListenAndServe(":3000", router))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

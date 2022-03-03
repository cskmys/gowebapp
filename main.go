// using templates to dynamically generate html content is going to simplify and make the design scalable

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
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
	t, err := template.ParseFiles("hello.gohtml") // "html.template::ParseFiles" is a variadic function, hence you can give a comma seperated list of file names
	if err != nil {
		panic(err)
	}

	data := User{ // this struct must be written as template expects an object rather than a direct string
		Name: "John Smith", // the field name must be "Name" coz template expects "{{.Name}}"
	}

	err = t.Execute(os.Stdout, data) // prints "<h1>Hello, John Smith</h1>" on stdout, "<h1>Hello, " and "</h1>" taken from template and
	// "John Smith" inserted into the template by "http.template::Execute"
	if err != nil {
		panic(err)
	}

	data.Name = "Jane Doe"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	http.ListenAndServe(":3000", router)
}

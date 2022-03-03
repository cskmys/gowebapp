// using "html/template" to prevent html injection attack

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
	t, err := template.ParseFiles("hello.gohtml") // "text.template::ParseFiles" is used instead of "html.template::ParseFiles"
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "<script>alert(\"hi\")</script>",
	}
	err = t.Execute(os.Stdout, data) // prints "<h1>Hello, &lt;script&gt;alert(&#34;hi&#34;)&lt;/script&gt;</h1>"
	// so when rendered by browser you'll see "Hello, <script>alert("hi")</script>"
	// so injection attack can happen
	// Hence, it is better to use context aware templating packages rather than plain "text/template" package
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

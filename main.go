// injection attacks can be used run nefarious code
// now we are using "text/template" package to illustrate what is an injection attack via html code
// simply using "html/template" is enough to prevent injection attack via html code

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
	"os"
	"text/template"
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
		Name: "<script>alert(\"hi\")</script>", // writing html code in text
	}
	err = t.Execute(os.Stdout, data) // prints "<h1>Hello, <script>alert("hi")</script></h1>"
	// when this text is served by server to browser, and the browser runs it, text "Hello, " and a pop-up saying "hi" can be seen on the browser
	// what if the "Name" field is supplied by an external agent and the string was a nefarious code
	// then harmful things can be done when the html is served by the server and run by the browser
	// this is injection attack
	if err != nil {
		panic(err)
	}

	// as a solution you can encode few characters such as "<", ">", "/" and "&"
	data.Name = html.EscapeString(data.Name) // after encoding "<h1>Hello, <script>alert("hi")</script></h1>" is converted into "<h1>Hello, &lt;script&gt;alert(&#34;hi&#34;)&lt;/script&gt;</h1>"
	err = t.Execute(os.Stdout, data)
	// Now when this text is served by server to browser, and the browser runs it, you'll just see "Hello, <script>alert("hi")</script>"
	// there is no pop-up like before
	// So, no nefarious things can be done when the string is encoded before being served to the browser
	if err != nil {
		panic(err)
	}

	// all the encoding logic to prevent html injection attack is already in "html/template" package, so you can simply use that instead of "text/template" package

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(custom404)

	http.ListenAndServe(":3000", router)
}

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound) // use the constant "http.StatusNotFound", don't write 404 directly
		// if "http.ResponseWriter::WriteHeader" is not used to set status code,
		// "http.ResponseWriter::Write" method that gets called from within "fmt::Fprint" will do "http.ResponseWriter::WriteHeader(http.StatusOK)" before writing the string
		// When "http.ResponseWriter::WriteHeader" is explicitly called to change status, you'll see the status change in response header(which you can see using Developer tools),
		// there won't be any visible change on the page
		fmt.Fprint(w, "<h1>Can't find the page you are looking for</h1><p>email me if you keep seeing this</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

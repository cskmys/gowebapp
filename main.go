package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my Awesome site!</h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" { // to make sure that "/contact" and "/contact/" gets served the same page
		fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Can't find the page you are looking for</h1><p>email me if you keep seeing this</p>")
	}
}

func main() {
	mux := &http.ServeMux{} // creating "mux" from default servemux
	mux.HandleFunc("abc/", handlerFunc)
	// if pattern was "/contact" then exact match "localhost:3000/contact" is expected to call "handlerFunc"
	// in this case, if there was no exact match d efault 404 page of servemux is returned
	// if pattern was "/contact/" then any URL that starts with "localhost:3000/contact" followed by "/" or "/<something>" gets routed to "handlerFunc"
	// hence, if there was no match 404 page written inside "handlerFunc" is returned
	// therefore, when "/" is used as a pattern, every request gets routed to "handlerFunc"
	/*
		// we can have multiple routing functions
		mux.HandleFunc("/contact/", handlerFunc2)
		mux.HandleFunc("/contact/me/", handlerFunc3)
		// now all URL that starts with "localhost:3000/contact/me" followed by "/" or  "/<something>" gets routed to "handlerFunc3"
		// all URL that starts with "localhost:3000/contact" followed by "/" or  "/<something>"(where <something> is not "me" or "me/" or "me/<something_else>") gets routed to "handlerFunc2"
		// all URL that starts with "localhost:3000" followed by "/" or  "/<something>"(where <something> is not "contact" or "contact/" or "contact/<something_else>") gets routed to "handlerFunc"
		// hence, the longest pattern gets matched first
	*/
	// as the pattern is "/", in the browser you can use the ip alias of the localhost i.e. 127.0.0.1 along with port no as "127.0.0.1:3000" to talk to server
	// if pattern was "localhost/", then in browser you cannot use the localhost alias, you'll have to do "localhost:3000"
	// however if pattern was "abc/", then in browser using "abc:3000" will lead to browser's 404(not server's 404) meaning browser is not able to talk to server
	// don't know why, maybe you'll need to register this application with OS or browser, so that browser will look for it to redirect port 3000 communications
	http.ListenAndServe(":3000", mux) // now we explicitly say what servemux should be used
}

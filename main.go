// here we replace the default servemux with a 3rd party servemux

// set "GOPATH" to somewhere else except where this file is
// create a new "go.mod" file by doing "go mod init"
// to install this "httprouter" package do "go get github.com/julienschmidt/httprouter"

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // ":name" regex matching is done with URL and the name field is available in "ps"
	// with default servemux we had to manually do the regex matching to extract name from url, here it is automatically done and available in "ps"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome %s!</h1>", ps.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/users/:name", Hello) // now, for all paths "localhost:3000/users/<name>", "Hello" gets called

	http.ListenAndServe(":3000", mux)
}

// to see all the request response content, go to:
// "Developer Tools", go to "Network" tab, there choose the type based on what you are looking for
// if you are looking for plain text type: "text/plain", choose "All" or "Other", for "text/html", choose "HTML" etc
// then select the browser request you want to examine, then choose "Headers" tab, there you can see
// "<HTML Method><URL>", "Response Headers", "Request Headers"
// under "Response Headers" you can see the key-value pairs of info served by server
// under "Request Headers" you'll see the key-value pairs of request made by the browser

package main

import (
	"fmt"
	"net/http"
)

func handlerFuncHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// web request/response are a bunch of key-value pairs and "Content-Type" is one of the keys
	// setting its value to "text/html" makes it clear that data sent is html
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
	// most browsers will directly take this as html content
	// but for some browsers it is not clear that you are sending html
	// hence you'll need to set few additional data to make it clear that it is html
}

func handlerFuncPlain(w http.ResponseWriter, r *http.Request) {
	// to display the string "<h1>Welcome to my awesome site</h1>"
	// make it clear that the text served is plain text
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func main() {
	http.HandleFunc("/", handlerFuncPlain)
	http.ListenAndServe(":3000", nil)
}

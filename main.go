package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<html>
	<head>
		<title>Lens Locked</title>
	</head>
	<body>
		<h1>Welcome to my pretty awesome site!</h1>
	</body>
</html>
`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch email me at <a href=\"mailto:bolaji@lenslocked.com\">bolaji@lenslocked.com</a>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router http.HandlerFunc = pathHandler

	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)

	//http.Handle("/", http.HandlerFunc(homeHandler))

	fmt.Println("Starting the server on port 4000")
	http.ListenAndServe(":4000", router)
}

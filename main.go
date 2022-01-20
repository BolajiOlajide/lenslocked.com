package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templatePath := filepath.Join("templates", "home.gohtml")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Printf("error executing the template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>Lens Locked | Contact</title>
	</head>
	<body>
		<h1>Contact page</h1>
		<p>To get in touch email me at <a href="mailto:bolaji@lenslocked.com">bolaji@lenslocked.com</a>
	</body>
</html>
`))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>Lens Locked | FAQs</title>
	</head>
	<body>
		<h1>Lens Locked FAQ</h1>
		<ul>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>Is there a free version?</h2>
			<p>Yes! We have a free trial for 30 days on any paid plans.</p>
		  </li>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>What are your support hours?</h2>
			<p>We have support staff answering emails 24/7, though response time may be a bit slower on weekends.</p>
		  </li>
		  <li>
			<input type="checkbox" checked>
			<i></i>
			<h2>How do I contact support?</h2>
			<p>Email us - <a href="support@lenslocked.com">support@lenslocked.com</a></p>
		  </li>
		</ul>
	</body>
</html>
`))
}

func getSingleResourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userID := chi.URLParam(r, "userID")
	fmt.Fprint(w, fmt.Sprintf(`<html>
	<head>
		<title>LensLocked | Single Resource</title>
	</head>
	<body>
		<h1>I am getting a single resource</h1>
		<p><b>userID</b>: %s<p>
	<body>
</html>
`, userID))
}

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Page Not Found", http.StatusNotFound)
//	}
//}

//type Router struct {}
//
//func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Page Not Found", http.StatusNotFound)
//	}
//}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.CleanPath, middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", homeHandler)
	r.Get("/user/{userID}", getSingleResourceHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	// not necessarily needed but okay to replicate former router structure
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)

	//http.Handle("/", http.HandlerFunc(homeHandler))

	fmt.Println("Starting the server on port 4000")
	http.ListenAndServe(":4000", r)
}

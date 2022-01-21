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

func executeTemplate(w http.ResponseWriter, templatePath string, substitution interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, substitution)
	if err != nil {
		log.Printf("error executing the template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, templatePath, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, templatePath, nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, templatePath, nil)
}

func getSingleResourceHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	isAdmin := chi.URLParam(r, "isAdmin") == "true"
	templatePath := filepath.Join("templates", "user", "{{id}}.gohtml")
	executeTemplate(w, templatePath, struct {
		UserID       string
		IsAdmin      bool
		Fruits       []string
		Salary       float32
		Measurements map[string]float32
		Age          int32
	}{
		UserID:  userID,
		IsAdmin: isAdmin,
		Fruits:  []string{"pineapple", "orange", "grapes"},
		Salary:  3130.28,
		Age:     12,
		Measurements: map[string]float32{
			"height": 34.45,
			"weight": 40.32,
		},
	})
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

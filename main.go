package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/BolajiOlajide/lenslocked.com/controllers"
	"github.com/BolajiOlajide/lenslocked.com/templates"
	"github.com/BolajiOlajide/lenslocked.com/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// type SomeType struct {
// 	Template views.Template
// }

// func (st SomeType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	st.Template.Execute(w, nil)
// }

func executeTemplate(w http.ResponseWriter, templatePath string, substitution interface{}) {
	template, err := views.Parse(templatePath)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	template.Execute(w, substitution)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	templatePath := filepath.Join("templates", "home.gohtml")
// 	executeTemplate(w, templatePath, nil)
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	templatePath := filepath.Join("templates", "contact.gohtml")
// 	executeTemplate(w, templatePath, nil)
// }

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	templatePath := filepath.Join("templates", "faq.gohtml")
// 	executeTemplate(w, templatePath, nil)
// }

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

	// parse the templates
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/user/{userID}", getSingleResourceHandler)

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

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

package controllers

import (
	"fmt"
	"net/http"
)

// User representing a lenslocked user information
type User struct {
	Template Template
}

// Template template for user controller
type Template struct {
	New ViewTemplate
}

// New creates a new user using the signup route
func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	// we need a view to render
	u.Template.New.Execute(w, data)
}

// Create backend implementation to actaually create the user
func (u User) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Fatalf("an error occurred while creqting user: %v", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println(r.PostForm.Get("email"), "<====")

	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(email, "<====", password)
	fmt.Fprintf(w, "Temporary Response")
}

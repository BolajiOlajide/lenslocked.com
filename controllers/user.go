package controllers

import (
	"fmt"
	"net/http"

	"github.com/BolajiOlajide/lenslocked.com/models"
)

// User representing a lenslocked user information
type User struct {
	Template    Template
	UserService *models.UserService
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

	user, err := u.UserService.Create(models.NewUser{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Created: %+v", user)
}

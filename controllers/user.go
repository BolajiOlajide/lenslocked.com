package controllers

import (
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
	// we need a view to render
	u.Template.New.Execute(w, nil)
}

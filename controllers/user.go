package controllers

import (
	"net/http"

	"github.com/BolajiOlajide/lenslocked.com/views"
)

// User representing a lenslocked user information
type User struct {
	Template Templates
}

type Templates struct {
	New views.Template
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	// we need a view to render
	u.Template.New.Execute(w, nil)
}

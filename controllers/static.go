package controllers

import (
	"net/http"

	"github.com/BolajiOlajide/lenslocked.com/views"
)

// StaticHandler takes a template and returns an handler func to be used to render the template
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

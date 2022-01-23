package controllers

import "net/http"

// ViewTemplate interface representing supported templates in project
type ViewTemplate interface {
	Execute(w http.ResponseWriter, data interface{})
}

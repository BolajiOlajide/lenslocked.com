package views

import (
	"html/template"
	"log"
	"net/http"
)

// Parse method used to parse html
func Parse(templatePath string) (Template, error) {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return Template{}, err
	}

	return Template{
		htmlTmpl: template,
	}, nil
}

// Template html template for use in lenslocked
type Template struct {
	htmlTmpl *template.Template
}

// Execute to actually execute the template with the substituion data
func (t Template) Execute(w http.ResponseWriter, substitution interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTmpl.Execute(w, substitution)
	if err != nil {
		log.Printf("error executing the template: %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}

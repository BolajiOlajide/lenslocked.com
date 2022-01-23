package controllers

import (
	"html/template"
	"net/http"
)

// StaticHandler takes a template and returns an handler func to be used to render the template
func StaticHandler(tpl ViewTemplate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

// FAQHandler handle faq but by passing in the faq data
func FAQHandler(tpl ViewTemplate) http.HandlerFunc {
	questions := []struct {
		Question string
		// using html here because it's not user provided, so it's safe
		Answer template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We have a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response time may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="support@lenslocked.com">support@lenslocked.com</a>`,
		},
		{
			Question: "Do you have an office?",
			Answer:   "Yes, but we are a remote-first company.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

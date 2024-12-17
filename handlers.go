package main

import (
	"net/http"
)

// If it starts with uppercase the function is public, else is private
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "about.page.tmpl")

}

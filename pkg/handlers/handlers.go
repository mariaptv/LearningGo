package handlers

import (
	"net/http"

	"github.com/mariaptv/go-course/pkg/render"
)

// If it starts with uppercase the function is public, else is private
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.html")

}

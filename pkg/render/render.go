package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error

	//check if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap{
		//need to create the template

	} else {
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s",t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		retur err
	}
}
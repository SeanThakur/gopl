package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.html")
}

func AboutHandler(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.html")
}

func renderTemplate(response http.ResponseWriter, templatePath string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + templatePath)
	if err != nil {
		println("something went wrong while parsing tempalte path", err.Error())
		return
	}
	err = parsedTemplate.Execute(response, nil)
	if err != nil {
		println("something went wrong while executing tempalte data", err.Error())
		return
	}
}

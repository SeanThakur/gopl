package handlers

import (
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/render"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "home.page.html")
}

func AboutHandler(response http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(response, "about.page.html")
}

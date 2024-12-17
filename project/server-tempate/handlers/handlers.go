package handlers

import (
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/models"
	"github.com/SeanThakur/gopl/project/server-tempate/render"
)

var Repo Repository

type Repository struct {
	App *config.APP_CONFIG
}

func NewRepo(a *config.APP_CONFIG) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = *r
}

func (Trepo *Repository) HomeHandler(response http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	Trepo.App.Session.Put(request.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(response, "home.page.html", &models.TemplateData{})
}

func (Trepo *Repository) AboutHandler(response http.ResponseWriter, request *http.Request) {
	remoteIP := Trepo.App.Session.GetString(request.Context(), "remote_ip")
	stringMapData := make(map[string]string)
	stringMapData["test"] = "data from server to about page"
	stringMapData["remote_ip"] = remoteIP
	render.RenderTemplate(response, "about.page.html", &models.TemplateData{
		StringMap: stringMapData,
	})
}

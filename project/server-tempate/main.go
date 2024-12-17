package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/handlers"
	"github.com/SeanThakur/gopl/project/server-tempate/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var session *scs.SessionManager
var app config.APP_CONFIG

func main() {
	app.IsProduction = false // set true for production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.IsProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	template, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("error encounted while creating template cache", err.Error())
		return
	}

	app.TemplateCache = template
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println("Starting port at ", PORT)
	serve := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	if err != nil {
		fmt.Println("Server crashed", err.Error())
		return
	}
}

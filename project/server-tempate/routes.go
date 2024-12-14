package main

import (
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.APP_CONFIG) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))

	return mux
}

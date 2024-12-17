package main

import (
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.APP_CONFIG) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(WriteConsole)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	return mux
}

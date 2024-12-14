package main

import (
	"fmt"
	"net/http"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/handlers"
	"github.com/SeanThakur/gopl/project/server-tempate/render"
)

const PORT = ":8080"

func main() {
	var app config.APP_CONFIG
	template, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("error encounted while creating template cache", err.Error())
		return
	}

	app.TemplateCache = template

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Starting port at ", PORT)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Server crashed", err.Error())
		return
	}
}

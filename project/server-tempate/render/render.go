package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/SeanThakur/gopl/project/server-tempate/config"
	"github.com/SeanThakur/gopl/project/server-tempate/models"
)

var app *config.APP_CONFIG

func NewTemplate(temp *config.APP_CONFIG) {
	app = temp
}

func RenderTemplate(response http.ResponseWriter, templatePath string, tempData *models.TemplateData) {
	var templateC map[string]*template.Template

	// create a template cache
	if app.UseCache {
		templateC = app.TemplateCache
	} else {
		templateC, _ = CreateTemplateCache()
	}

	// get requested template from the cache
	templateG, ok := templateC[templatePath]
	if !ok {
		fmt.Println("did not get tempalte cache")
		return
	}

	buffer := new(bytes.Buffer)
	err := templateG.Execute(buffer, tempData)
	if err != nil {
		fmt.Println("Error encounered while executing template buffer", err.Error())
		return
	}

	// render the template
	_, err = buffer.WriteTo(response)
	if err != nil {
		fmt.Println("Error encounered while writing and rendering tempalte", err.Error())
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all the page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return templateCache, err
	}

	// parse all the pages one by one and include template layout file in all the pages
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matchesLayout, err := filepath.Glob("./templates/*.template.html") // layout file
		if err != nil {
			return templateCache, err
		}

		if len(matchesLayout) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.template.html")
			if err != nil {
				return templateCache, err
			}
		}

		// add this template to the cache
		templateCache[name] = templateSet
	}

	return templateCache, nil
}

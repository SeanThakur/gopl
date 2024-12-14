package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(response http.ResponseWriter, templatePath string) {
	// create a template cache
	templateC, err := CreateTemplateCache()
	if err != nil {
		fmt.Println("Error encounered while creating template cache", err.Error())
		return
	}
	// get requested template from the cache
	templateG, ok := templateC[templatePath]
	if !ok {
		fmt.Println("did not get tempalte cache")
		return
	}

	buffer := new(bytes.Buffer)
	err = templateG.Execute(buffer, nil)
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

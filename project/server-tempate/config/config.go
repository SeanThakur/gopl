package config

import "html/template"

type APP_CONFIG struct {
	TemplateCache map[string]*template.Template
}

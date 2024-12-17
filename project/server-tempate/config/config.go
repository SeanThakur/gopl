package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type APP_CONFIG struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	Session       *scs.SessionManager
	IsProduction  bool
}

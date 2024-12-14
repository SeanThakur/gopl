package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRFToken string
	Success   string
	Warning   string
	Error     string
}

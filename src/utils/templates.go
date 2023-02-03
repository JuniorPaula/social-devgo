package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func Render(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}

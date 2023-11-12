package templates

import (
	"html/template"
	"net/http"
)

var Templates = template.Must(template.ParseGlob("templates/*.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := Templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"html/template"
	"net/http"

	"github.com/tonitienda/progressive-enhancement-.git/pkg/tasks"
)

var (
	htmlTemplate = template.Must(template.ParseFiles("templates/html.html"))
)

func RenderHtml(w http.ResponseWriter) {

	htmlTemplate.Execute(w, tasks.GetTasks())
}

package handlers

import (
	"html/template"
	"net/http"

	"github.com/tonitienda/progressive-enhancement-.git/pkg/tasks"
)

var (
	reactTemplate = template.Must(template.ParseFiles("templates/react.html"))
)

func RenderReact(w http.ResponseWriter) {
	reactTemplate.Execute(w, tasks.GetTasks())
}

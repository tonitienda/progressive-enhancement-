package handlers

import (
	"html/template"
	"net/http"

	"github.com/tonitienda/progressive-enhancement-.git/pkg/tasks"
)

var (
	htmxTemplate = template.Must(template.ParseFiles("templates/htmx.html"))
)

func RenderHtmx(w http.ResponseWriter) {
	htmxTemplate.Execute(w, tasks.GetTasks())
}

// func RenderHtmxTaskList(w http.ResponseWriter, tasks []tasks.Task) {

// 	w.Header().Set("Content-Type", "text/html")

// 	htmxTemplate.Execute(w, "task-list", tasks)
// }

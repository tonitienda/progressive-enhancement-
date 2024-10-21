package main

import (
	"encoding/json"
	"net/http"

	"github.com/tonitienda/progressive-enhancement-.git/pkg/handlers"
	"github.com/tonitienda/progressive-enhancement-.git/pkg/tasks"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/tasks", tasksHandler)

	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	app := r.URL.Query().Get("app")
	if app == "react" {
		handlers.RenderReact(w)
	} else if app == "htmx" {
		handlers.RenderHtmx(w)
	} else {
		handlers.RenderHtml(w)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	switch r.Method {
	// case http.MethodGet:
	// 	handleGetTasks(w, contentType)
	case http.MethodPost:
		handlePostTasks(w, r, contentType)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// func handleGetTasks(w http.ResponseWriter, contentType string) {
// 	tasksList := tasks.GetTasks()
// 	if contentType == "application/json" {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(tasksList)
// 	} else {
// 		w.Header().Set("Content-Type", "text/html")
// 		htmxTemplate.ExecuteTemplate(w, "task-list", tasksList)
// 	}
// }

func getTaskFromRequest(w http.ResponseWriter, r *http.Request, contentType string) (tasks.Task, bool) {
	if contentType == "application/json" {
		var newTask tasks.Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return tasks.Task{}, false
		}

		return newTask, true
	}

	if contentType == "application/x-www-form-urlencoded" {
		taskText := r.FormValue("task")
		task := tasks.Task{Text: taskText, Completed: false}

		return task, true
	}

	return tasks.Task{}, false
}

func handlePostTasks(w http.ResponseWriter, r *http.Request, contentType string) {
	task, ok := getTaskFromRequest(w, r, contentType)

	if !ok {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newId := tasks.AddTask(task)

	w.Header().Set("Location", "/tasks/"+newId)

	// if r.Header.Get("HX-Request") != "" {
	// 	handlers.RenderHtmxTaskList(w, tasks.GetTasks())
	// } else {
	// 	w.WriteHeader(http.StatusCreated)
	// }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HTML Templates (react.html and index.html)
/*
   react.html:
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>React TODO List</title>
       <script src="https://unpkg.com/react@17/umd/react.development.js"></script>
       <script src="https://unpkg.com/react-dom@17/umd/react-dom.development.js"></script>
       <script src="/static/app.js"></script>
   </head>
   <body>
       <div id="root"></div>
   </body>
   </html>

   index.html:
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>TODO List</title>
       <script src="https://unpkg.com/htmx.org@1.6.1"></script>
   </head>
   <body>
       <h1>TODO List</h1>
       <div>
           <form hx-post="/tasks" hx-target="#task-list" hx-swap="innerHTML">
               <input type="text" name="task" placeholder="Add a new task" required>
               <button type="submit">Add Task</button>
           </form>
       </div>
       <ul id="task-list">
           {{ range $index, $task := . }}
           <li>
               <span hx-post="/tasks/toggle" hx-target="#task-list" hx-swap="innerHTML" hx-vals='{"index": "{{ $index }}"}'>{{ if $task.Completed }}<s>{{ $task.Text }}</s>{{ else }}{{ $task.Text }}{{ end }}</span>
               <button hx-post="/tasks/delete" hx-target="#task-list" hx-swap="innerHTML" hx-vals='{"index": "{{ $index }}"}'>Delete</button>
           </li>
           {{ end }}
       </ul>
   </body>
   </html>
*/

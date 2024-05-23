package task

import (
	"fmt"
	"net/http"

	task "github.com/teooliver/kanban/internal/repository/task"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// here we read from the request context and fetch out `"user"` key set in
	// the MyMiddleware example above.
	// task := r.Context().Value("user").(string)
	// r.Body.Read(p []byte)

	// respond to the client
	w.Write([]byte(fmt.Sprintf("Task %s", "task")))
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	task.ListTasks(db)

	w.Write([]byte(fmt.Sprintf("Task %s", "task")))
}

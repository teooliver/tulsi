package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teooliver/kanban/internal/repository/task"
)

type taskService interface {
	CreateTask(ctx context.Context, req task.TaskForCreate) (*task.Task, error)
	ListTasks(ctx context.Context) ([]task.Task, error)
}

type Handler struct {
	svc taskService
}

func New(svc taskService) Handler {
	return Handler{
		svc: svc,
	}
}

type ListTaskResponse struct {
	Tasks []task.Task `json:"tasks"`
}

func (h Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.svc.ListTasks(ctx)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong")))
	}
	taskResponse := ListTaskResponse{
		Tasks: tasks,
	}

	jsonTasks, err := json.Marshal(taskResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(jsonTasks))
}

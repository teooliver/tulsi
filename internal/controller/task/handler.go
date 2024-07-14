package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/task"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type taskService interface {
	ListAllTasks(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[task.Task], error)
	CreateTask(ctx context.Context, task task.TaskForCreate) error
	DeleteTask(ctx context.Context, taskID string) error
	UpdateTask(ctx context.Context, taskID string, updatedTask task.TaskForUpdate) error
}

type Handler struct {
	service taskService
}

func New(service taskService) Handler {
	return Handler{
		service: service,
	}
}

// type ListTasksInput struct {
// 	Token  string  `in:"header=Authorization;omitempty"`
// 	Page   int     `in:"query=page;default=1"`
// 	Size   int     `in:"query=size;default=20"`
// 	Search *string `in:"query=search;omitempty"`
// }

type ListTaskResponse struct {
	Tasks postgresutils.Page[task.Task] `json:"tasks"`
}

func (h Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := r.Context().Value(httpin.Input).(*postgresutils.PageRequest)
	slog.Info("HTTPin Input: ", input)

	tasks, err := h.service.ListAllTasks(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER => Something went wrong: %v\n", err)))
	}
	taskResponse := ListTaskResponse{
		Tasks: tasks,
	}

	jsonTasks, err := json.Marshal(taskResponse.Tasks)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER MARSHAL => Something went wrong: %v\n", err)))
	}

	w.Write([]byte(jsonTasks))
}

func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var taskToCreate task.TaskForCreate
	err := json.NewDecoder(r.Body).Decode(&taskToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("Task for CREATE %+v\n", "taskToCreate", taskToCreate)

	err = h.service.CreateTask(ctx, taskToCreate)

}

func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := chi.URLParam(r, "id")
	err := h.service.DeleteTask(ctx, taskID)

	if err != nil {
		// Should return Error Not Found and 404
		slog.Info("TaskID %+v\n", taskID, taskID)
	}

}

func (h Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := chi.URLParam(r, "id")
	var taskToUpdate task.TaskForUpdate
	err := json.NewDecoder(r.Body).Decode(&taskToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateTask(ctx, taskID, taskToUpdate)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
	}

}

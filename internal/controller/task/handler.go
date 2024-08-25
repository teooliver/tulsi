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
	CreateTask(ctx context.Context, task task.TaskForCreate) (string, error)
	DeleteTask(ctx context.Context, taskID string) (string, error)
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

	tasks, err := h.service.ListAllTasks(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER => Something went wrong: %v\n", err)))
		return
	}
	taskResponse := ListTaskResponse{
		Tasks: tasks,
	}

	jsonTasks, err := json.Marshal(taskResponse.Tasks)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
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

	id, err := h.service.CreateTask(ctx, taskToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := chi.URLParam(r, "id")
	id, err := h.service.DeleteTask(ctx, taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Delete Task - Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
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
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(taskID))

}

package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/task"
)

type taskService interface {
	ListAllTasks(ctx context.Context) ([]task.Task, error)
	CreateTask(ctx context.Context, task task.TaskForCreate) error
	DeleteTask(ctx context.Context, taskID string) error
	UpdateTask(ctx context.Context, taskID string, updatedTask task.TaskForUpdate) error
	InsertMultipleTasks(ctx context.Context) error
}

type Handler struct {
	service taskService
}

func New(service taskService) Handler {
	return Handler{
		service: service,
	}
}

type ListTaskResponse struct {
	Tasks []task.Task `json:"tasks"`
}

func (h Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.service.ListAllTasks(ctx)
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

func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var taskToCreate task.TaskForCreate
	err := json.NewDecoder(r.Body).Decode(&taskToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("Task for CREATE %+v\n", taskToCreate, taskToCreate)

	err = h.service.CreateTask(ctx, taskToCreate)

}

func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := chi.URLParam(r, "id")

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	slog.Info("TaskID %+v\n", taskID, taskID)

	err := h.service.DeleteTask(ctx, taskID)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
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

	slog.Info("TaskID %+v\n", taskID, taskID)
	slog.Info("TaskForUpdate %+s\n", taskToUpdate)

	err = h.service.UpdateTask(ctx, taskID, taskToUpdate)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
	}

}

func (h Handler) SeedTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	slog.Info("HELLO")

	err := h.service.InsertMultipleTasks(ctx)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
	}

}

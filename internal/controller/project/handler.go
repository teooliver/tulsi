package project

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/column"
	"github.com/teooliver/kanban/internal/repository/project"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type projecService interface {
	ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[project.Project], error)
	GetProjectColumns(ctx context.Context, projectId string) ([]column.Column, error)
	ArquiveProject(ctx context.Context, projectID string) (string, error)
	CreateProject(ctx context.Context, project project.CreateProjectRequest) (string, error)
}

type Handler struct {
	service projecService
}

func New(service projecService) Handler {
	return Handler{
		service: service,
	}
}

type ListProjectResponse struct {
	Projects postgresutils.Page[project.Project] `json:"projects"`
}

func (h Handler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := r.Context().Value(httpin.Input).(*postgresutils.PageRequest)

	projects, err := h.service.ListAllProjects(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Project HANDLER => Something went wrong: %v\n", err)))
		return
	}
	projectResponse := ListProjectResponse{
		Projects: projects,
	}

	jsonProjects, err := json.Marshal(projectResponse.Projects)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("List Projects HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonProjects))
}

func (h Handler) GetProjectColumns(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := chi.URLParam(r, "id")

	columns, err := h.service.GetProjectColumns(ctx, projectID)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Get Project Columns HANDLER => Something went wrong: %v\n", err)))
		return
	}

	jsonColumns, err := json.Marshal(columns)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("Get Project Columns HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonColumns))
}

func (h Handler) ArquiveProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectID := chi.URLParam(r, "id")
	id, err := h.service.ArquiveProject(ctx, projectID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Arquive Project - Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}
func (h Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var projectRequest project.CreateProjectRequest
	err := json.NewDecoder(r.Body).Decode(&projectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("Project for CREATE %+v\n", "projectToCreate", projectRequest)

	id, err := h.service.CreateProject(ctx, projectRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Create Project - Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

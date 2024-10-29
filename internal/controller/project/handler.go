package project

import (
	"context"
	"encoding/json"
	"fmt"
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
		w.Write([]byte(fmt.Sprintf("Project HANDLER MARSHAL => Something went wrong: %v\n", err)))
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
		w.Write([]byte(fmt.Sprintf("Project HANDLER => Something went wrong: %v\n", err)))
		return
	}

	jsonColumns, err := json.Marshal(columns)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("Project HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonColumns))
}

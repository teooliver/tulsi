package column

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teooliver/tulsi/internal/repository/column"
)

type columnService interface {
	CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error)
	GetColumnsByProjectID(ctx context.Context, projectID string) ([]column.Column, error)
}

type Handler struct {
	service columnService
}

func New(service columnService) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var columnRequest column.ColumnForCreate
	err := json.NewDecoder(r.Body).Decode(&columnRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("GOT HERE - HANDLER CREATE COLUMN %+v\n", "ColumnToCreate", columnRequest)

	id, err := h.service.CreateColumn(ctx, columnRequest)
	// slog.Info("ERROR", "ColumnToCreate", err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Create Column - Something went wrong: %v\n", err)))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func (h Handler) GetColumnsByProjectID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID := chi.URLParam(r, "projectID")

	columns, err := h.service.GetColumnsByProjectID(ctx, projectID)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Get Columns by ProjectID HANDLER => Something went wrong: %v\n", err)))
		return
	}

	jsonColumns, err := json.Marshal(columns)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("Get Columns by ProjectID HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonColumns))
}

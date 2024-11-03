package column

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/teooliver/kanban/internal/repository/column"
)

type columnService interface {
	CreateColumn(ctx context.Context, column column.ColumnForCreate) (string, error)
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

	slog.Info("Column for CREATE %+v\n", "ColumnToCreate", columnRequest)

	id, err := h.service.CreateColumn(ctx, columnRequest)
	slog.Info("ERROR", "ColumnToCreate", err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Create Column - Something went wrong: %v\n", err)))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}
